import axios from 'axios'

export const apiInstance = axios.create({
  baseURL: 'http://localhost:8080/',
});

export const cloudinaryInstance = axios.create({
  baseURL: 'http://api.cloudinary.com/v1_1/dadowgksf'
})

export default apiInstance