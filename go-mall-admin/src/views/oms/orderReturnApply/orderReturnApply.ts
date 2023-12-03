const defaultStatusOptions = [
  {
    label: "待处理",
    value: 0
  },
  {
    label: "退货中",
    value: 1
  },
  {
    label: "已完成",
    value: 2
  },
  {
    label: "已拒绝",
    value: 3
  }
];
export const formatStatus = (status: number) => {
  for (let i = 0; i < defaultStatusOptions.length; i++) {
    if (status === defaultStatusOptions[i].value) {
      return defaultStatusOptions[i].label;
    }
  }
};
