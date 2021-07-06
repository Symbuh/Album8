import axios from 'axios'

export const apiInstance = axios.create({
  baseURL: 'http://localhost:3000',
});

// apiInstance.defaults.headers.post['Content-Type'] = 'application/json'
apiInstance.defaults.headers.post['Access-Control-Allow-Origin'] = '*'

export const cloudinaryInstance = axios.create({
  baseURL: 'http://api.cloudinary.com/v1_1/dadowgksf'
})

export default apiInstance