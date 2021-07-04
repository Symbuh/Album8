import React, { FC, useState, useEffect, useCallback } from 'react'
import apiInstance from './../axiosConfig'
import Image from './Image'
import MainImageView from './SelectedImageDisplay/MainImageView'
import UploadFile from './UploadFile'
import SearchByTagToggle from './SearchByTag/SearchByTagToggle'

const Carousel: FC = () => {

  const [updateImage, setUpdateImage] = useState(false)
  const [images, setImages]: any[] = useState([]);
  const [allTags, setAllTags]: any[] = useState([])
  const [{
          id,
          name,
          url,
          description,
          tags
        }, setSelectedImage] = useState({
          id: '',
          name: '',
          url: '',
          description: '',
          tags: []
        });

  useEffect(() => {
    getImages()
    getTags()
  }, [])

  const handleClick = (id: string, url: string, name: string, description: string, tags: any) => {
    setSelectedImage({
      id: id,
      url: url,
      name: name,
      description: description,
      tags: tags
    })
  }

  const getImages = () => {
    apiInstance.get('/api/image')
    .then((response: any) => {
      setImages(response.data)
    })
    .catch((err: any) => {
      console.log(err);
    })
  }

  const getTags = () => {
    apiInstance.get('/api/tags')
    .then((response: any) => {
      setAllTags(response.data)
    })
    .catch((err: any) => {
      console.log(err)
    })
  }

  const updateFunction = () => {
    getImages()
    getTags()
    setSelectedImage({id: '', name: '', url: '', description: '', tags: []})
  }

  return (
    <div>
      <div>
        <UploadFile updateFunction={updateFunction}/>
      </div>
      <div id="searchByTagToggleContainer">
        <SearchByTagToggle tags={allTags} setImagesByTag={setImages}/>
      </div>
      {
        images.map((image: any) => {
          const {id, url, name, description, tags} = image
          return (
            <div onClick={() => handleClick( id, url, name, description, tags )}>
              <Image
                id={id}
                url={url}
                name={name}
                description={description}
                tags={tags}
              />
            </div>
          )
        })
      }
      <div>
        <MainImageView
          id={id}
          name={name}
          url={url}
          description={description}
          tags={tags}
          updateFunction={updateFunction}
        />
      </div>
    </div>
  )
}

export default Carousel