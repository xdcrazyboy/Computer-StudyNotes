# React入门笔记

## 为什么要用React

### 组件化
1. 可复用，class类，然后render{return一个html 字符串}
   >还得去获取里面的DOM，添加实践，就需要一个：一个函数 createDOMFromString ，你往这个函数传入 HTML 字符串，但是它会把相应的 DOM 元素返回给你。




## JSX语法
1. XML基本语法
   - 定义标签时，只允许被一个标签包裹：最外层必须时一个大标签包着，不能出现两个同等级的标签作为最外层。
   - **标签一定要闭合。** html自闭和的标签也要改成左右闭合。

2. 元素类型
   - 小写字母： DOM元素
   - 大写首字母： 组件元素
   >命名空间，解决组件相同名称冲突，比如： <MUI.RaisedButton label="Default" />

3. 注释
