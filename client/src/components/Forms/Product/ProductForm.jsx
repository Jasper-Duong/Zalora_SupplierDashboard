/* eslint-disable react-hooks/exhaustive-deps */
import { Form, Input, InputNumber, Switch } from "antd";
import React, { useEffect } from "react";
import FormBtns from "../FormBtns";
import ProductFormCancelBtn from "./ProductFormCancelBtn";
import FormContainer from "../../Container/FormContainer";
import { useForm } from "antd/es/form/Form";

export default function ProductForm({ product, onFinish, submitBtnText }) {
  const [form] = useForm()
  useEffect(() => {
    form.setFieldsValue({ ...product })
  }, [product])
  return (
    <FormContainer>
      <Form
        onFinish={onFinish}
        layout="vertical"
        form={form}
        initialValues={{status: true}}
        labelAlign="left"
      >
        <Form.Item
          label="Name"
          name="name"
          rules={[{ required: true, message: "Product Name is required" }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          label="Brand Name"
          name="brand"
          rules={[{ required: true, message: "Brand Name is required" }]}
        >
          <Input />
        </Form.Item>
        {product && <><Form.Item
          label="SKU"
          name="sku"
        >
          <Input disabled/>
        </Form.Item>
        <Form.Item
          label="Stock"
          name="stock"
          rules={[{ required: true, message: "Product Stock is required" }]}
        >
          <InputNumber min={0} disabled/>
        </Form.Item></>}
        <Form.Item
          label="Color"
          name="color"
          rules={[{ required: true, message: "Product Color is required" }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          label="Size"
          name="size"
          rules={[{ required: true, message: "Product Size is required" }]}
        >
          <Input />
        </Form.Item>
        <Form.Item label="Status" name="status" valuePropName="checked">
          <Switch checkedChildren="Active" unCheckedChildren="Inactive"/>
        </Form.Item>
        {/* <Form.Item
          label="Suppliers"
          name="suppliers"
          rules={[
            {
              required: true,
              message: "Product must has at least 1 supplier",
              type: "array",
            },
          ]}
        >
          <Select mode="multiple" options={supplierOptions} />
        </Form.Item> */}
        <FormBtns
          CancelBtn={<ProductFormCancelBtn />}
          submitBtnText={submitBtnText}
        />
      </Form>
    </FormContainer>
  );
}
