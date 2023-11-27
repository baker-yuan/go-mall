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
    console.log("setup", props.modelValue);

    const formatValue = (value: string) => {
      return value.replace(/\n/g,",");
    };

    const formattedValue = ref(props.modelValue.replace(/,/g,"\n"));

    watch(() => props.modelValue, (newValue, oldValue) => {
      console.log("watch", newValue, oldValue);
      if (newValue) {
        formattedValue.value = newValue.replace(/,/g,"\n");
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
