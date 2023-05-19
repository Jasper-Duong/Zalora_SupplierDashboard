import { Col, Form, Input, Row, Select } from "antd";
import { useForm } from "antd/es/form/Form";

export default function AddRowAddress({ onFinish, AddConfirmBtn }) {
  const [form] = useForm();
  const onSelectChange = (value) => {
    console.log({ value });
    form.setFieldValue("type", value);
  };
  return (
    <Form onFinish={onFinish} style={{ width: "100%" }} form={form}>
      <Form.Item hidden name={"id"} />
      <Row gutter={16}>
        <Col span={7}>
          <Form.Item label="Type" name={"type"} rules={[{ required: true }]}>
            <Select onChange={onSelectChange}>
              <Select.Option value="office">Office</Select.Option>
              <Select.Option value="warehouse">Warehouse</Select.Option>
            </Select>
          </Form.Item>
        </Col>
        <Col span={12}>
          <Form.Item
            label={"Address Info"}
            name={"address_info"}
            rules={[{ required: true }]}
          >
            <Input />
          </Form.Item>
        </Col>
        <Col span={5}>
          <AddConfirmBtn />
        </Col>
      </Row>
    </Form>
  );
}
