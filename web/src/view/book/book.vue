<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">   
        <el-form-item label="小说标题">
          <el-input placeholder="搜索条件" v-model="searchInfo.bookTitle"></el-input>
        </el-form-item>    
        <el-form-item label="小说编号">
          <el-input placeholder="搜索条件" v-model="searchInfo.bookId"></el-input>
        </el-form-item>    
        <el-form-item label="作者">
          <el-input placeholder="搜索条件" v-model="searchInfo.bookWriter"></el-input>
        </el-form-item>    
        <el-form-item label="小说类型">
          <el-input placeholder="搜索条件" v-model="searchInfo.bookType"></el-input>
        </el-form-item>    
        <el-form-item label="章节编号">
          <el-input placeholder="搜索条件" v-model="searchInfo.sectionNum"></el-input>
        </el-form-item>    
        <el-form-item label="章节标题">
          <el-input placeholder="搜索条件" v-model="searchInfo.sectionTitle"></el-input>
        </el-form-item>      
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增book表</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
              </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="55"></el-table-column>
    <el-table-column label="日期" width="180">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>
    
    <el-table-column label="原网站" prop="srcWeb" width="120"></el-table-column> 
    
    <el-table-column label="小说标题" prop="bookTitle" width="120"></el-table-column> 
    
    <el-table-column label="作者" prop="bookWriter" width="120"></el-table-column> 
    
    <el-table-column label="小说类型" prop="bookType" width="120"></el-table-column> 
    
    <el-table-column label="章节编号" prop="sectionNum" width="120"></el-table-column> 
    
    <el-table-column label="章节标题" prop="sectionTitle" width="220"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="updateBook(scope.row)" size="small" type="primary">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteBook(scope.row)">确定</el-button>
            </div>
            <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
          </el-popover>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
          <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="80px"
            label-position="left">
            <el-row type="flex" justify="start" align="top" gutter="15">
              <el-form-item label="原网站" prop="srcWeb">
                <el-input v-model="formData.srcWeb" placeholder="请输入原网站" clearable :style="{width: '100%'}">
                </el-input>
              </el-form-item>
              <el-form-item label="原域名" prop="srcHost">
                <el-input v-model="formData.srcHost" placeholder="请输入原域名" clearable :style="{width: '100%'}">
                </el-input>
              </el-form-item>
              <el-form-item label="原网址" prop="srcUrl">
                <el-input v-model="formData.srcUrl" placeholder="请输入原网址" clearable :style="{width: '100%'}">
                </el-input>
              </el-form-item>
            </el-row>
            <el-row type="flex" justify="start" align="top" gutter="15">
              <el-form-item label="小说编号" prop="bookId">
                <el-input v-model="formData.bookId" placeholder="请输入小说编号" clearable :style="{width: '100%'}">
                </el-input>
              </el-form-item>
              <el-form-item label="小说标题" prop="bookTitle">
                <el-input v-model="formData.bookTitle" placeholder="请输入小说标题" clearable :style="{width: '100%'}">
                </el-input>
              </el-form-item>
              <el-form-item label="小说类型" prop="bookType">
                <el-input v-model="formData.bookType" placeholder="请输入小说类型" clearable :style="{width: '100%'}">
                </el-input>
              </el-form-item>
              <el-form-item label="作者" prop="bookWriter">
                <el-input v-model="formData.bookWriter" placeholder="请输入作者" clearable :style="{width: '100%'}">
                </el-input>
              </el-form-item>
            </el-row>
            <el-row gutter="15">
              <el-form-item label="章节编号" prop="sectionNum">
                <el-input v-model="formData.sectionNum" placeholder="请输入章节编号" clearable :style="{width: '100%'}">
                </el-input>
              </el-form-item>
              <el-form-item label="章节标题" prop="sectionTitle">
                <el-input v-model="formData.sectionTitle" placeholder="请输入章节标题" clearable :style="{width: '100%'}">
                </el-input>
              </el-form-item>
              <el-form-item label="备注" prop="remarks">
                <el-input v-model="formData.remarks" placeholder="请输入备注" clearable :style="{width: '100%'}">
                </el-input>
              </el-form-item>
            </el-row>
            <el-form-item label="章节内容" prop="sectionContent">
              <el-input v-model="formData.sectionContent" type="textarea" placeholder="请输入章节内容"
                :autosize="{minRows: 4, maxRows: 14}" :style="{width: '100%'}"></el-input>
            </el-form-item>
          </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
    createBook,
    deleteBook,
    deleteBookByIds,
    updateBook,
    findBook,
    getBookList
} from "@/api/book";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "Book",
  mixins: [infoList],
  data() {
    return {
      listApi: getBookList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        srcWeb:null,srcHost:null,srcUrl:null,bookTitle:null,bookId:null,bookWriter:null,bookType:null,sectionNum:null,sectionTitle:null,sectionContent:null,remarks:null,
      }
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" :"否";
      } else {
        return "";
      }
    }
  },
  methods: {
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10               
        this.getTableData()
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      async onDelete() {
        const ids = []
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await deleteBookByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateBook(row) {
      const res = await findBook({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.rebook;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        
          srcWeb:null,
          srcHost:null,
          srcUrl:null,
          bookTitle:null,
          bookId:null,
          bookWriter:null,
          bookType:null,
          sectionNum:null,
          sectionTitle:null,
          sectionContent:null,
          remarks:null,
      };
    },
    async deleteBook(row) {
      this.visible = false;
      const res = await deleteBook({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createBook(this.formData);
          break;
        case "update":
          res = await updateBook(this.formData);
          break;
        default:
          res = await createBook(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  created() {
    this.getTableData();
  }
};
</script>

<style>
</style>