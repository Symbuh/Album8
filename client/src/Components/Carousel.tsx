import React, { FC, useState, useEffect } from 'react'
import axios from 'axios'
import apiInstance from './../axiosConfig'
import Image from './Image'
import MainImageView from './SelectedImageDisplay/MainImageView'
import UploadFile from './UploadFile'


const Carousel: FC = () => {
  /*
    Get all images from axios and save them in state,

    map through these images and display them to the page.
  */

    const [images, setImages] = useState([]);
    const [{id, name, url, description, tags}, setSelectedImage] = useState({id: '', name: '', url: '', description: '', tags: ''});

    const handleClick: any = (id: string, url: string, name: string, description: string, tags: any) => {
      setSelectedImage({
        id: id,
        url: url,
        name: name,
        description: description,
        tags: tags
      })
    }

    useEffect(() => {
      getImages()
    }, [])

    const getImages = () => {
      apiInstance.get('/api/image')
      .then((response: any) => {
        setImages(response.data)
      })
      .catch((err: any) => {
        console.log(err);
      })
    }

    return (
      <div>
        <div>
          <UploadFile updateFunction={getImages}/>
        </div>
        {
          images.map((image) => {
            const {id, url, name, description, tags} = image
            return (
              <div onClick={handleClick( id, url, name, description, tags)}>
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
            updateFunction={getImages}
          />
        </div>
      </div>
    )
}

export default Carousel