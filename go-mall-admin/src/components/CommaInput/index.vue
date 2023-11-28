<!--
一、Vue 3 Composition API
-->
<!--
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
  // 声明"update:modelValue"事件
  emits: ["update:modelValue"],
  setup(props, { emit }) {
    const convertValue = (value: string) => {
      return value.replace(/,/g, "\n");
    };
    const formatValue = (value: string) => {
      return value.replace(/\n/g, ",");
    };

    const formattedValue = ref(convertValue(props.modelValue));

    watch(
      () => props.modelValue,
      (newValue, oldValue) => {
        console.log("watch props.modelValue", newValue, oldValue);
        if (newValue) {
          formattedValue.value = convertValue(newValue);
        }
      }
    );

    watch(formattedValue, (newValue, oldValue) => {
      console.log("watch formattedValue", newValue, oldValue);
      emit("update:modelValue", formatValue(newValue));
    });

    return {
      formattedValue
    };
  }
};
</script>
-->

<!--
二、更简洁的 Composition API 写法 （语法糖）
-->
<template>
  <el-input :autosize="true" type="textarea" v-model="formattedValue"></el-input>
</template>

<script setup lang="ts">
import { ref, watchEffect } from "vue";
import { ElInput } from "element-plus";

const props = defineProps({
  modelValue: {
    type: String,
    default: ""
  }
});

const emit = defineEmits(["update:modelValue"]);

const convertValue = (value: string) => {
  return value.replace(/,/g, "\n");
};
const formatValue = (value: string) => {
  return value.replace(/\n/g, ",");
};

let formattedValue = ref(convertValue(props.modelValue));

watchEffect(() => {
  formattedValue.value = convertValue(props.modelValue);
});

watchEffect(() => {
  emit("update:modelValue", formatValue(formattedValue.value));
});
</script>

<!--
三、Options API（Vue 2 的写法）
-->
<!--
<template>
  <el-input :autosize="true" type="textarea" v-model="formattedValue"></el-input>
</template>

<script lang="ts">
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
  data() {
    return {
      formattedValue: this.convertValue(this.modelValue)
    };
  },
  methods: {
    convertValue(value: string) {
      return value.replace(/,/g, "\n");
    },
    formatValue(value: string) {
      return value.replace(/\n/g, ",");
    },
    updateValue(value: string) {
      this.formattedValue = this.convertValue(value);
      this.$emit("update:modelValue", this.formatValue(this.formattedValue));
    }
  },
  watch: {
    modelValue: {
      handler(newValue) {
        this.updateValue(newValue);
      },
      immediate: true
    },
    formattedValue(newValue) {
      this.$emit("update:modelValue", this.formatValue(newValue));
    }
  }
};
</script>
-->
