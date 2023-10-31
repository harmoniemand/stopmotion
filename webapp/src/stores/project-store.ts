import { defineStore } from "pinia"
import { ref } from "vue"
import Project from "../types/Project"
import { computed } from "@vue/reactivity"
import Image from "../types/Image"
import axios from "axios"

export const useProjectStore = defineStore('project', () => {
    console.log("creating project store")

    const baseURL = "http://localhost:8080"
    // const baseURL = "https://stopmotion.sfzprojekt.de"

    const project = ref<Project | null>(null)

    const hasImages = computed(() => {
        if (project.value == null) return false

        return project.value?.images.length > 0
    })
    const lastImage = computed(() => {
        return project.value?.images.slice().reverse()[0]
    })

    const AddImage = (image: Image) => {
        if (project.value == null) {
            return
        }

        project.value.images.push(image)


        const splitDataURI = image.data.split(',')
        const byteString = splitDataURI[0].indexOf('base64') >= 0 ? atob(splitDataURI[1]) : decodeURI(splitDataURI[1])
        const mimeString = splitDataURI[0].split(':')[1].split(';')[0]

        const ia = new Uint8Array(byteString.length)
        for (let i = 0; i < byteString.length; i++)
            ia[i] = byteString.charCodeAt(i)

        let b = new Blob([ia], { type: mimeString })

        let formData = new FormData();
        formData.append("project_id", project.value.id);
        formData.append("id", image.Id);
        formData.append("image", b, image.Id + ".png");


        axios.post(baseURL + '/api/v1/images', formData, {
            headers: {
                'Content-Type': 'multipart/form-data'
            }
        });
    }

    const RemoveImage = (image: Image) => {
        if (project.value == null) {
            return
        }

        project.value.images = project.value.images.filter(i => i != image)
    }

    function LoadImageAsBase64(pid: string, iid: string) {
        console.log("loading image", pid, iid)

        let img = document.createElement('img');
        img.src = baseURL + '/api/v1/images/' + pid + '/' + iid;
        img.crossOrigin = "anonymous"
        img.onload = function () {
            let canvas = document.createElement('canvas');
            canvas.width = img.width;
            canvas.height = img.height;

            let ctx = canvas.getContext('2d');
            ctx?.drawImage(img, 0, 0);

            let dataURL = canvas.toDataURL('image/png');

            // let image = new Image({id: iid, data: dataURL})
            // console.log("image loaded", image)

            let i = project.value?.images.findIndex(i => i.Id == iid)
            if (i == undefined) {
                console.warn("image not found", iid)
                return
            }

            if (project.value == null) {
                console.warn("project not found")
                return
            }

            project.value.images[i].data = dataURL
        };

    }

    function LoadProject() {
        console.log("looking for project_id")
        const project_id = localStorage.getItem('project_id')

        if (project_id) {
            console.log("project_id found, fetching project data", project_id)
            fetch(baseURL + '/api/v1/projects/' + project_id)
                .then(response => response.json())
                .then(data => {
                    project.value = new Project(data)
                    console.log("project data fetched", project.value)

                    project.value.images.forEach(i => {
                        LoadImageAsBase64(project_id, i.Id)
                    })
                })
                .catch(error => {
                    console.log(error)
                })
        }
    }

    function CreateProject(p: Project) {

        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(p)
        };

        fetch(baseURL + '/api/v1/projects', requestOptions)
            .then(response => response.json())
            .then(data => {
                project.value = new Project(data)
                localStorage.setItem('project_id', data.id)
            })
            .catch(error => {
                console.log(error)
            })
    }

    return {
        project,
        hasImages,
        LoadProject,
        CreateProject,
        AddImage,
        RemoveImage

    }
})
