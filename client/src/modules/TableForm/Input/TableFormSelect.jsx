import { Select } from "antd";
import useService from "../../../hooks/useService";

const onSearch = (value) => {
  console.log("search:", value);
};

const TableFormSelect = ({ form, service, placeholder }) => {
  const { data } = useService({
    service: service,
  });
  const options = data?.map((item) => ({
    value: item.id,
    label: item.name,
  }));
  const onChange = (value) => {
    console.log(`selected ${value}`);
    const { label } = options.find((item) => item.value === value);
    if (form) {
      form.setFieldValue("id", value);
      form.setFieldValue("name", label);
    }
  };
  return (
    <Select
      showSearch
      placeholder={placeholder}
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
