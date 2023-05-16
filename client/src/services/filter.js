import axios from '../config/axios'

async function getFilterOptions(options) {
    return axios.get(`${options}`)
}

export {
    getFilterOptions
}