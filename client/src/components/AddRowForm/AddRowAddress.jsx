import { Col, Form, Input, Row, Select } from "antd";
import { useForm } from "antd/es/form/Form";

export default function AddRowAddress({ onFinish, AddConfirmBtn }) {
  const [form] = useForm();
  const onSelectChange = () => {
  }
  return (
    <Form onFinish={onFinish} style={{width: "100%"}} form={form}>
      <Form.Item hidden name={"id"} />
      <Row gutter={16}>
        <Col span={7}>
          <Form.Item>
            <Select defaultValue={"office"} onChange={onSelectChange}>
              <Select.Option value="office">Office</Select.Option>
              <Select.Option value="warehouse">Warehouse</Select.Option>
            </Select>
          </Form.Item>
        </Col>
        <Col span={12}>
          <Form.Item label={"Address Info"} name={"address_info"}>
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
