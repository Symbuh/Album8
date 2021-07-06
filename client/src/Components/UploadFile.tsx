import React, {FC, useEffect, useState} from 'react'
import { isIndexedAccessTypeNode } from 'typescript'
import { cloudinaryInstance } from './../axiosConfig'
import UploadModal from './UploadDetailsModal/UploadModal'


interface Props {
  updateFunction: () => void
}

const UploadFile: FC<Props> = ({ updateFunction }) => {
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
      setURL(res.data.url)
    })
    .catch((err) => {
      console.error('Failed to upload image!')
    })
  }

  useEffect(() => {
    if (selectedImage !== '') {
      sendToCloudinary()
    }
  }, [selectedImage])

  return (
    <div className='imageUploadContainer'>
      <input type='file' onChange={handleChange} />
      {
        newURL !== '' &&
        <UploadModal url={newURL} updateFunction={updateFunction}/>
      }
    </div>
  )
}

export default UploadFile