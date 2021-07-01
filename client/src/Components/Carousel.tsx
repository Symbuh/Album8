import React, { FC, useState, useEffect } from 'react'
import axios from 'axios'
import apiInstance from './../axiosConfig'
import Image from './Image'

const Carousel: FC = () => {
  /*
    Get all images from axios and save them in state,

    map through these images and display them to the page.
  */

    const [images, setImages] = useState([]);

    useEffect(() => {
      apiInstance.get('/api/image')
      .then((response: any) => {
        setImages(response.data)
      })
      .catch((err: any) => {
        console.log(err);
      })
    }, [])

    return (
      <div>
        {
          images.map((image) => {
            const {id, url, name, description, image_tags} = image
            return (
              <div>
                <Image
                  id={id}
                  url={url}
                  name={name}
                  description={description}
                  tags={image_tags}
                />
              </div>
            )
          })
        }
      </div>
    )
}

export default Carousel