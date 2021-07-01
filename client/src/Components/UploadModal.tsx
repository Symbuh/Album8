import React, { FC, useState } from 'react'
import Modal from 'react-modal'
import { addSyntheticTrailingComment, setConstantValue } from 'typescript'

interface Props {
  url: string
}

const UploadModal: FC<Props> = ( { url }) => {
  const [modalIsOpen, setIsOpen] = useState(false)
  const [imageObject, setImg] = useState({name: '', description: '', url: url, tag: ''})
  const [tags, setTags] = useState([])



  /*
    It would be nice to save the form data as an object so that we can easily toss it into our
    API

    We want an image object with name url desc and tags

    I think that we may need a seperate state variable to push our tags
  */
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
          name='tag'
          onChange={handleChange}
        />
        <button onClick={addTag}>Add Tag</button>
      </Modal>
    </div>
  )
}

export default UploadModal