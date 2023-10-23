<template>
    <div class="toolbar">
        <button class="logo">
            <img src="./../assets/Logo.png" />
        </button>

        <button class="toolbar_item" @click="onButtonNewPressed">Neu</button>
    </div>
</template>


<script setup lang="ts">
const onButtonNewPressed = () => {
    console.debug("button new pressed")
}
</script>


<style scoped>
.toolbar {
    width: 100%;
    background-color: #222222;
    height: 46px;
    /* border: solid 1px blue; */
    display: flex;
}

.toolbar_item {
    color: hsla(160, 100%, 37%, 1);
    background-color: rgba(0, 0, 0, 0);
    border: solid 1px red;
    cursor: pointer;
    padding: 0px 20px;
}

.logo {
    border: none;
    background-color: rgba(0, 0, 0, 0);
    padding: 5px;
    margin-right: 25px;
}

.logo img {
    height: 36px;
    background-color: rgba(0, 0, 0, 0);
}
</style>
