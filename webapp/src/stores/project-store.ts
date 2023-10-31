import Project from "../types/Project"
import Image from "../types/Image"
import axios from "axios"
import { ref } from "vue";


class ProjectStore {

    private static instance: ProjectStore;


    baseURL = "http://localhost:8080"
    // baseURL = "https://stopmotion.sfzprojekt.de"
    Project = ref<Project | null>(null)

    private constructor() {
        console.log("creating project store")
    }

    public static getInstance(): ProjectStore {
        if (!ProjectStore.instance) {
            ProjectStore.instance = new ProjectStore();
        }

        return ProjectStore.instance;
    }

    AddImage(image: Image) {
        if (this.Project.value == null) {
            return
        }

        this.Project.value.images.push(image)


        const splitDataURI = image.data.split(',')
        const byteString = splitDataURI[0].indexOf('base64') >= 0 ? atob(splitDataURI[1]) : decodeURI(splitDataURI[1])
        const mimeString = splitDataURI[0].split(':')[1].split(';')[0]

        const ia = new Uint8Array(byteString.length)
        for (let i = 0; i < byteString.length; i++)
            ia[i] = byteString.charCodeAt(i)

        let b = new Blob([ia], { type: mimeString })

        let formData = new FormData();
        formData.append("project_id", this.Project.value.id);
        formData.append("id", image.Id);
        formData.append("image", b, image.Id + ".png");


        axios.post(this.baseURL + '/api/v1/images', formData, {
            headers: {
                'Content-Type': 'multipart/form-data'
            }
        });
    }

    RemoveImage(image: Image) {
        if (this.Project.value == null) {
            return
        }

        this.Project.value.images = this.Project.value.images.filter(i => i != image)
    }

    LoadProject() {
        console.log("looking for project_id")
        const project_id = localStorage.getItem('project_id')

        if (project_id) {
            console.log("project_id found, fetching project data", project_id)
            fetch(this.baseURL + '/api/v1/projects/' + project_id)
                .then(response => response.json())
                .then(data => {
                    this.Project.value = new Project(data)
                    console.log("project data fetched", this.Project)

                    this.Project.value.images.forEach(i => {
                        this.LoadImageAsBase64(project_id, i.Id)
                    })
                })
                .catch(error => {
                    console.log(error)
                })
        }
    }

    LoadImageAsBase64(pid: string, iid: string) {
        console.log("loading image", pid, iid)

        let img = document.createElement('img');
        img.src = this.baseURL + '/api/v1/images/' + pid + '/' + iid;
        img.crossOrigin = "anonymous"
        img.onload = () => {
            let canvas = document.createElement('canvas');
            canvas.width = img.width;
            canvas.height = img.height;

            let ctx = canvas.getContext('2d');
            ctx?.drawImage(img, 0, 0);

            let dataURL = canvas.toDataURL('image/png');

            if (this.Project.value == null) {
                console.warn("project not found")
                return
            }

            let i = this.Project.value.images.findIndex(i => i.Id == iid)
            if (i == undefined) {
                console.warn("image not found", iid)
                return
            }

            this.Project.value.images[i].data = dataURL
        };
    }

    CreateProject(p: Project) {

        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(p)
        };

        fetch(this.baseURL + '/api/v1/projects', requestOptions)
            .then(response => response.json())
            .then(data => {
                this.Project.value = new Project(data)
                localStorage.setItem('project_id', data.id)
            })
            .catch(error => {
                console.log(error)
            })
    }
}


export default ProjectStore