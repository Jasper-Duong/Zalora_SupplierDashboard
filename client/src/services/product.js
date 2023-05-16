import { request } from '../config/axios'

function getProductByIdApi(id) {
    return request.get(`products/${id}`)
}

function updateProductByIdApi(id, product) {
    return request.put(`products/${id}`, JSON.stringify(product))
}

function getProductStocks(id) {
    return request.get(`products/${id}/stocks`)
}
export {
    getProductByIdApi,
    updateProductByIdApi
}