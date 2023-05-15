import { Form, Input, InputNumber, Switch } from "antd";
import React from "react";
import { regexPatterns } from "../../../constants/forms";
import FormBtns from "../FormBtns";
import SupplierFormCancelBtn from "./SupplierFormCancelBtn";
import FormContainer from "../../Container/FormContainer";

export default function SupplierForm({ supplier, onFinish, submitBtnText }) {
  return (
    <FormContainer>
        <Form
          layout="vertical"
          onFinish={onFinish}
          initialValues={supplier ? { ...supplier, stock: 10 } : { stock: 10 }}
        >
          <Form.Item
            label="Name"
            name="name"
            rules={[{ required: true, message: "Supplier Name is required" }]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            label="Email"
            name="email"
            rules={[
              {
                require: true,
                message: "Supplier Email is required",
              },
              { pattern: regexPatterns.email, message: "Invalid Email" },
            ]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            label="Contact Number"
            name="contactNumber"
            rules={[
              {
                required: true,
                message: "Supplier Contact Number is required",
              },
              {
                pattern: regexPatterns.phoneNumber,
                message: "Invalid Contact Number",
              },
            ]}
          >
            <Input />
          </Form.Item>
          <Form.Item label="Stock" name="stock">
            <InputNumber disabled min={0} />
          </Form.Item>
          <Form.Item label="Status" name="status" valuePropName="checked">
            <Switch checkedChildren="Active" unCheckedChildren="Inactive" />
          </Form.Item>
          <FormBtns
            CancelBtn={<SupplierFormCancelBtn />}
            submitBtnText={submitBtnText}
          />
        </Form>
        </FormContainer>
  );
}
