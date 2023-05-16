import { request } from "../config/axios";

async function updateSupplierApi() {}
async function getSupplierByIdApi(id) {
  return request({ method: "GET", url: `/suppliers/${id}` });
}
