import axios from '../config/axios'

function getProductByIdApi(id) {
    return axios.get(`products/${id}`)
}

function updateProductByIdApi(id, product) {
    return axios.put(`products/${id}`, JSON.stringify(product))
}

function getProductStocks(id) {
    return axios.get(`products/${id}/stocks`)
}
export {
    getProductByIdApi,
    updateProductByIdApi
}