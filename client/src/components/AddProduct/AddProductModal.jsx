import CustomModal from "../Modals/CustomModal";
import AddProductForm from "../Forms/Product/AddProductForm";

export default function AddProductModal({ isOpen }) {
  return (
    <CustomModal isOpen={isOpen} title={"Add Product"} Child={AddProductForm} />
  );
}
