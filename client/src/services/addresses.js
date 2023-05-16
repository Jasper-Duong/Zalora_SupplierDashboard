import { request } from "../config/axios";

async function getAddressesApi(id) {
  return request({ method: "GET", url: `/suppliers/${id}/addresses` });
}

export { getAddressesApi };
