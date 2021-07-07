import React, { FC, useState, useEffect } from 'react'
import Modal from 'react-modal'
import { apiInstance } from '../../axiosConfig'
import axios from 'axios'
import SelectedTags from './SelectedTags'

interface Props {
  url: string,
  updateFunction: () => void
}

const UploadModal: FC<Props> = ( { url, updateFunction }) => {
  const [modalIsOpen, setIsOpen] = useState(false)
  const [formComplete, setFormComplete] = useState(false)
  const [imageObject, setImg] = useState({
    name: '', description: '', url: url, tags: ''
  })
  const [tags, setTags] = useState([])

  useEffect(() => {
    const { name, description } = imageObject
    if (name !== '' && description !== '') {
      if(tags.length > 0) {
        setFormComplete(true)
      }
    }
  }, [imageObject, tags])

  // have to break this into a smaller file
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
    const existingTags: any = tags as any
    if (imageObject.tags !== '') {
      if (!existingTags.includes(imageObject.tags)) {
        existingTags.push(imageObject.tags)
        setTags(existingTags)

        setImg({
          ...imageObject,
          tags: ''
        })
      }
    }
  }

  const handleClick = () => {
    sendToAPI()
    updateFunction()
  }

  const sendToAPI = () => {
    console.log('calling update function when sending to API')
    if (formComplete) {
      const requestBody: any = imageObject
      requestBody.tags = tags
      const config: any = {
        method: 'post',
        url: 'http://localhost:8080/api/newimage',
        headers: {
          'Content-Type': 'application/json'
        },
        data: JSON.stringify(requestBody)
      };

      axios(config)
      .then(function (response) {
        console.log(JSON.stringify(response.data));
      })
      .catch(function (error) {
        console.log(error);
      });
    }
    closeModal()
    updateFunction()
    setImg({
      name: '',
      description: '',
      url: '',
      tags: ''
    })
  }

  return (
    <div>
      <button id="uploadImageButton" onClick={openModal}>Upload Image</button>
      <Modal
        isOpen={modalIsOpen}
        onRequestClose={closeModal}
      >
        <h2>Provide Image Details</h2>
        <h4>Please fill out all fields</h4>
        <input
          type='text'
          placeholder='Photo Name'
          name='name'
          value={imageObject.name}
          onChange={handleChange}/>
        <input
          type='text'
          placeholder='Description'
          name='description'
          value={imageObject.description}
          onChange={handleChange}
        />
        <input
          type='text'
          placeholder='Tag'
          name='tags'
          value={imageObject.tags}
          onChange={handleChange}
        />
        <button onClick={addTag}>Add Tag</button>
        <SelectedTags tags={tags} />
        {
          formComplete && <button onClick={handleClick}>Submit</button>
        }
      </Modal>
    </div>
  )
}

export default UploadModal