const formLayout = { labelCol: { span: 8 }, wrapperCol: { span: 16 } };

const regexPatterns = {
  // eslint-disable-next-line
  phoneNumber: /^[+]*[(]{0,1}[0-9]{1,4}[)]{0,1}[-\s\./0-9]*$/g,
  // eslint-disable-next-line
  email: /^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/g,
};

export { formLayout, regexPatterns };
