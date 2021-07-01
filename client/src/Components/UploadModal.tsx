import React, { FC, useState, useEffect } from 'react'
import Modal from 'react-modal'
import { apiInstance } from './../axiosConfig'

interface Props {
  url: string
}

const UploadModal: FC<Props> = ( { url }) => {
  const [modalIsOpen, setIsOpen] = useState(false)
  const [formComplete, setFormComplete] = useState(false)
  const [imageObject, setImg] = useState({
    name: '', description: '', url: url, tags: ''
  })
  const [tags, setTags] = useState([])

  useEffect(() => {
    const { name, description } = imageObject
    console.log(`Calling useEffect name: ${name}, description: ${description}`)
    if (name !== '' && description !== '') {
      if(tags.length > 0) {
        console.log('oh jeez man')
        setFormComplete(true)
      }
    }
  }, [imageObject, tags])


  const openModal = () => {
    setIsOpen(true)
  }

  const closeModal = () => {
    setIsOpen(false)
  }

  const handleChange = (e: any) => {
    setImg({
      ...imageObject,
      [e.target.name]: e.target.value
    })
  }

  const addTag = () => {
    let existingTags: any = tags as any
    if (imageObject.tags !== '') {
      if (!existingTags.includes(imageObject.tags)) {
        existingTags.push(imageObject.tags)
        setTags(existingTags)
      }
    }
  }

  const sendToAPI = () => {
    if (formComplete) {
      const requestBody: any = imageObject
      requestBody.tags = tags
      apiInstance.post('/api/newimage', requestBody)
      .then(res => {
        console.log(res)
      })
      .catch(err => {
        console.log(err)
      })
    }
    closeModal()
  }

  return (
    <div>
      <button onClick={openModal}>Upload Image</button>
      <Modal
        isOpen={modalIsOpen}
        onRequestClose={closeModal}
      >
        <h2>Provide Image Details</h2>
        <input
          type='text'
          placeholder='Photo Name'
          name='name'
          onChange={handleChange}/>
        <input
          type='text'
          placeholder='Description'
          name='description'
          onChange={handleChange}
        />
        <input
          type='text'
          placeholder='Tag'
          name='tags'
          onChange={handleChange}
        />
        <button onClick={addTag}>Add Tag</button>
        {
          formComplete && <button onClick={sendToAPI}>Submit</button>
        }
      </Modal>
    </div>
  )
}

export default UploadModal