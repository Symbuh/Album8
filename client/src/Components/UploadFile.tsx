import { time } from 'console'
import React, {FC, useState} from 'react'
import { isIndexedAccessTypeNode } from 'typescript'
import { cloudinaryInstance, apiInstance } from './../axiosConfig'
import UploadModal from './UploadModal'

const UploadFile: FC = () => {
  const [selectedImage, setImage] = useState('')
  const [newURL, setURL] = useState('')

  const handleChange = (event: any) => {
    setImage(event.target.files[0])
  }

  const sendToCloudinary = () => {
    const data = new FormData()
    console.log(selectedImage)
    data.append('file', selectedImage)
    data.append('upload_preset', 'sj5ltk6e')
    // data.append('signature', )
    cloudinaryInstance.post('/image/upload', data)
    .then((res: any) => {
      console.log(res)
      setURL(res.secure_url)
    })
    .catch((err) => {
      console.error('Failed to upload image!')
    })
  }

  const sendToAPI = () => {
    apiInstance.get('/')
    .then((res: any) => {
      console.log(res)
    })
    .catch((err: any)=> console.log(err))
  }

  // const sendFile = () => {
  //   const accept = ['image/png']
  //   if (selectedImage) {
  //     if (accept.indexOf(selectedImage.mediaType) > -1) {

  //     }
  //   }

  // }
  return (
    <div className='imageUploadConatainer'>
      <input type='file' onChange={handleChange} />
      <button onClick={sendToCloudinary}>Upload File</button>
      {
        newURL !== undefined &&
        <UploadModal url={newURL}/>
      }
    </div>
  )
}

export default UploadFile