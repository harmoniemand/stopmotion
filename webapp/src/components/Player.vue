<template>
    <div class="camera">
      <div class="camera-video">
        <div style="position: relative;">
          <img :src="images[currentImage].BASE64" />
        </div>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  
  import { ref } from 'vue'
  import ImageService from '../services/image.service'
  

  const imageService = ImageService.getInstance()
  const fps = ref(10);

  const images = imageService.Images
  const currentImage = ref(0)

  setInterval(() => {
    currentImage.value++;
    currentImage.value = currentImage.value % images.value.length;
  }, 1000 / fps.value)

  
  
  </script>
  
  <style scoped>
  .camera {
    width: 100%;
    height: 100%;
    background-color: #212121;
  
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
  
    position: relative;
  }
  
  .camera .camera-error-container {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
  
    display: flex;
    align-items: center;
    justify-content: center;
  
  }
  
  .camera .camera-error-container .camera-error {
  
    width: 250px;
    height: 250px;
  
    border: solid 1px hsla(17, 100%, 37%, 1);
    background-color: hsla(17, 100%, 37%, 0.2);
  
    display: flex;
    align-items: center;
    justify-content: center;
  }
  </style>
  