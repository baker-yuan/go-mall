<template>
  <el-input :autosize="true" type="textarea" v-model="formattedValue"></el-input>
</template>

<script lang="ts" name="CommaInput">
import { ref, watch } from "vue";
import { ElInput } from "element-plus";

export default {
  name: "CommaInput",
  components: {
    ElInput
  },
  props: {
    modelValue: {
      type: String,
      default: ""
    }
  },
  setup(props, { emit }) {
    const convertValue = (value: string) => {
      return value.replace(/,/g,"\n");
    };
    const formatValue = (value: string) => {
      return value.replace(/\n/g,",");
    };

    const formattedValue = ref(convertValue(props.modelValue));

    watch(() => props.modelValue, (newValue, oldValue) => {
      console.log("watch", newValue, oldValue);
      if (newValue) {
        formattedValue.value = convertValue(newValue);
      }
    })

    watch(formattedValue, (newValue, oldValue) => {
      console.log("watch", newValue, oldValue);
      emit('update:modelValue', formatValue(newValue))
    });

    return {
      formattedValue
    }
  }
}
</script>
