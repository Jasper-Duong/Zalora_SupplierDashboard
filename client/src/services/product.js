import { request } from '../config/axios'

function getProductByIdApi(id) {
    return request.get(`products/${id}`)
}

async function addProductByIdApi(product) {
    return request({ method: "POST", url: `products/`, data: JSON.stringify(product)})
}

function updateProductByIdApi(id, product) {
    return request.put(`products/${id}`, JSON.stringify(product))
}

export {
    getProductByIdApi,
    addProductByIdApi,
    updateProductByIdApi
}