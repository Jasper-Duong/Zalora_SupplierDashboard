import { Button, Modal } from "antd";
import styled from "styled-components";
import React, { useState } from "react";

const AntdModal = ({
  openModalBtnText,
  title,
  BodyComponent,
  ShowModalBtn,
}) => {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const showModal = () => {
    setIsModalOpen(true);
  };
  const closeModal = () => {
    setIsModalOpen(false);
  };
  return (
    <Styled>
      {ShowModalBtn ? (
        <ShowModalBtn showModal={showModal} />
      ) : (
        <Button type="primary" onClick={showModal}>
          {openModalBtnText || "Open Modal"}
        </Button>
      )}

      <Modal
        destroyOnClose
        title={title || "Basic Modal"}
        open={isModalOpen}
        onCancel={closeModal}
        footer={null}
      >
        <BodyComponent closeModal={closeModal} />
      </Modal>
    </Styled>
  );
};
export default AntdModal;

const Styled = styled.div``;
