// 轮播位置：0->PC首页轮播；1->app首页轮播
export const formatType = (type: number) => {
  if (type === 1) {
    return "APP首页轮播";
  } else {
    return "PC首页轮播";
  }
};

// 广告位置
export const typeOptions = [
  {
    label: "PC首页轮播",
    value: 0
  },
  {
    label: "APP首页轮播",
    value: 1
  }
];
