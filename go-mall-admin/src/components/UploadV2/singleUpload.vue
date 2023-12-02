<template>
  <div>
    <el-upload
      :action="useOss ? ossUploadUrl : minioUploadUrl"
      :data="useOss ? dataObj : null"
      list-type="picture"
      :multiple="false"
      :show-file-list="showFileList"
      :file-list="fileList"
      :before-upload="beforeUpload"
      :on-remove="handleRemove"
      :on-success="handleUploadSuccess"
      :on-preview="handlePreview"
    >
      <el-button size="small" type="primary">点击上传</el-button>
      <template #tip>
        <div class="el-upload__tip">只能上传jpg/png文件，且不超过10MB</div>
      </template>
    </el-upload>
    <el-dialog v-model:visible="dialogVisible">
      <img width="100%" :src="fileList[0].url" alt="" />
    </el-dialog>
  </div>
</template>
<script>
import { policy } from "@/api/modules/oss.ts";

export default {
  name: "SingleUpload",
  props: {
    value: String
  },
  data() {
    return {
      dataObj: {
        policy: "",
        signature: "",
        key: "",
        ossaccessKeyId: "",
        dir: "",
        host: ""
        // callback:'',
      },
      dialogVisible: false,
      useOss: false, //使用oss->true;使用MinIO->false
      ossUploadUrl: "http://macro-oss.oss-cn-shenzhen.aliyuncs.com",
      minioUploadUrl: "http://localhost:8080/minio/upload"
    };
  },
  computed: {
    imageUrl() {
      return this.value;
    },
    imageName() {
      if (this.value != null && this.value !== "") {
        return this.value.substr(this.value.lastIndexOf("/") + 1);
      } else {
        return null;
      }
    },
    fileList() {
      return [
        {
          name: this.imageName,
          url: this.imageUrl
        }
      ];
    },
    showFileList: {
      get: function () {
        return this.value !== null && this.value !== "" && this.value !== undefined;
      },
      // eslint-disable-next-line @typescript-eslint/no-unused-vars,@typescript-eslint/no-empty-function
      set: function (newValue) {}
    }
  },
  methods: {
    emitInput(val) {
      this.$emit("input", val);
    },
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    handleRemove(file, fileList) {
      this.emitInput("");
    },
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    handlePreview(file) {
      this.dialogVisible = true;
    },
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    beforeUpload(file) {
      // eslint-disable-next-line @typescript-eslint/no-this-alias
      let _self = this;
      if (!this.useOss) {
        //不使用oss不需要获取策略
        return true;
      }
      return new Promise((resolve, reject) => {
        policy()
          .then(response => {
            _self.dataObj.policy = response.data.policy;
            _self.dataObj.signature = response.data.signature;
            _self.dataObj.ossaccessKeyId = response.data.accessKeyId;
            _self.dataObj.key = response.data.dir + "/${filename}";
            _self.dataObj.dir = response.data.dir;
            _self.dataObj.host = response.data.host;
            // _self.dataObj.callback = response.data.callback;
            resolve(true);
          })
          .catch(err => {
            console.log(err);
            reject(false);
          });
      });
    },
    handleUploadSuccess(res, file) {
      this.showFileList = true;
      this.fileList.pop();
      let url = this.dataObj.host + "/" + this.dataObj.dir + "/" + file.name;
      if (!this.useOss) {
        //不使用oss直接获取图片路径
        url = res.data.url;
      }
      this.fileList.push({ name: file.name, url: url });
      this.emitInput(this.fileList[0].url);
    }
  }
};
</script>
<style></style>
