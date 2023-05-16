import { Select } from "antd";

const onSearch = (value) => {
  console.log("search:", value);
};

const options = [...Array(10).keys()].map((idx) => ({
  value: 100 + idx,
  label: `New Product ${idx + 1}`,
}));

const TableFormSelect = ({ form }) => {
  const onChange = (value) => {
    console.log(`selected ${value}`);
    const { label } = options.find((item) => item.value === value);
    if (form) {
      form.setFieldValue('id', value);
      form.setFieldValue("name", label);
    }
  };
  return (
    <Select
      showSearch
      placeholder="Select a product"
      optionFilterProp="children"
      onChange={onChange}
      onSearch={onSearch}
      filterOption={(input, option) =>
        (option?.label ?? "").toLowerCase().includes(input.toLowerCase())
      }
      options={options}
    />
  );
};
export default TableFormSelect;
