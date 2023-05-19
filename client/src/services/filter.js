import {request} from '../config/axios'

async function getFilterOptions(options) {
    return request.get(`${options}`)
}

export {
    getFilterOptions
}