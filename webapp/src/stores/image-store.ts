import { defineStore } from "pinia"
import Image from "../types/Image"
import { ref } from "vue"

export const useTodos = defineStore('todos', {
    state: () => ({
        images: ref<Image[]>([]),
    }),
    getters: {
        reversedImages(state) {
            return state.images.slice().reverse()
        },
    },
    actions: {
        addImage(image: Image) {
            this.images.push(image)
        },
    },
})
