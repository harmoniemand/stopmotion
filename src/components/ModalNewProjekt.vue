<template>
    <div class="overlay">
        <div class="overlay-inner">
            <div class="header">{{ currentStage.title }}</div>
            <div class="content">


                <section v-if="currentStage.id == 1">
                    <div class="text">
                        <p>Starten wir ein neues Projekt zusammen. Heute wollen wir einen StopMotion Film drehen!</p>
                        <p>Dazu werden wir mit LEGO eine Welt bauen. In dieser Welt spielt unser Film.</p>
                        <p>Beim StopMotion Film werden immer wieder Bilder aufgenommen, bei dem Kleinigkeiten geändert
                            werden. Aus allen Bildern ergibt sich ein Film.</p>
                    </div>
                    <div class="img">
                        <img src="./../assets/welcome.gif" alt="welcome" />
                    </div>
                </section>

                <section v-if="currentStage.id == 2">
                    <div class="text">
                        <p>Wie LEGO funktioniert muss ich euch glaube ich nicht erklären!</p>
                        <p>Aus vielen Steinen könnt ihr euch eine Welt bauen. Außerdem könnt ihr alles Papier und andere
                            Material nutzen.</p>
                        <p>Als Schauspielerinnen und Schauspieler treten dann die Legomännchen auf.</p>
                    </div>
                    <div class="img">
                        <img src="./../assets/lego_hero.jpg" alt="welcome" />
                    </div>
                </section>


                <section v-if="currentStage.id == 3">
                    <div class="text">
                        <p>Jetzt müsst ihr Bild für Bild eine Kleinigkeit ändern und immer wieder ein neues Bild aufnehmen.
                        </p>
                        <p>Mit jedem Bild wächst euer Film. In Kinofilmen werden 25 Bilder pro Sekunde gemacht. Ganz so
                            viele brauchen wir aber nicht.</p>
                        <p>Wenn ihr Hilfe braucht, fragt uns gerne.</p>
                    </div>
                    <div class="img">
                        <img src="./../assets/action.png" alt="welcome" />
                    </div>
                </section>


                <section v-if="currentStage.id == 4">
                    <div class="text">
                        <p>Für euren Film benötigen wir einen Name. Wie soll euer Film heisen?</p>
                        <p>
                            <input type="text" v-model="data.name" />
                        </p>
                    </div>
                    <div class="img">
                        <img src="./../assets/plakat.jpg" alt="welcome" style="background-size: contain;" />
                        <div class="movie-name">
                            {{ data.name }}
                        </div>
                    </div>
                </section>

                <section v-if="currentStage.id == 5">
                    <div class="text">
                        <p>Jetzt brauchen wir noch eine kleine Beschreibung. Eventuell könnt ihr in zwei Sätzen eure
                            Filmidee beschreiben?</p>
                        <p>
                            <textarea v-model="data.description" />
                        </p>
                    </div>
                    <div class="img">
                        <img src="./../assets/plakat.jpg" alt="welcome" style="background-size: contain;" />
                        <div class="movie-name">
                            {{ data.name }}
                        </div>
                    </div>
                </section>

                <section v-if="currentStage.id == 6">
                    <div class="text">
                        <p>Als letztes brauchen wir noch eure Vornamen.</p>
                        <p>
                            <textarea v-model="data.authors" />
                        </p>
                    </div>
                    <div class="img">
                        <img src="./../assets/plakat.jpg" alt="welcome" style="background-size: contain;" />
                        <div class="movie-name">
                            {{ data.name }}
                        </div>
                    </div>
                </section>

            </div>

            <div class="footer">
                <button v-if="currentStage.id < 6" class="btn btn-next" @click="next">Weiter</button>
                <button v-if="currentStage.id == 6" class="btn btn-next" @click="ready">Los Gehts</button>
                <button class="btn btn-back" :disabled="currentStage.id == 1" @click="back">Zurück</button>
            </div>
        </div>
    </div>
</template>


<script setup lang="ts">

import { computed, ref } from 'vue'

const data = ref<any>({
    name: 'the lego movie',
    description: 'Ein Film mit LEGO',
    authors: "joni"
})

const currentStageId = ref(1)
const stages = ref<any[]>([
    { id: 1, title: "Lass uns beginnen!" },
    { id: 2, title: "Lego!" },
    { id: 3, title: "Action!" },
    { id: 4, title: "Name" },
    { id: 5, title: "Beschreibung" },
    { id: 6, title: "Wer seid ihr?" },
])
const currentStage = computed(() => {
    return stages.value.find(s => s.id == currentStageId.value)
})

const emit = defineEmits(['ready'])

const back = () => {
    if (currentStageId.value > 1)
        currentStageId.value--;
}
const next = () => {
    currentStageId.value++;
}

const ready = () => {
    emit('ready', data.value)
}

</script>



<style scoped>
@font-face {
    font-family: 'bangers';
    src: url(./../assets/Bangers/Bangers-Regular.ttf);
}

.overlay {
    position: absolute;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;

    background-color: rgba(0, 0, 0, 0.5);

    display: flex;
    justify-content: center;
    align-items: center;
}

.overlay .overlay-inner {
    height: 70vh;
    width: 70vw;

    background-color: white;
    color: #535bf2;

    border-radius: 10px;

    display: flex;
    flex-direction: column;
    justify-content: space-between;
}

.overlay .overlay-inner .header {
    padding: 1rem;
    background-color: hsla(160, 100%, 37%, 1);
    color: white;
    font-size: 25px;
    font-weight: bold;
    border-radius: 10px 10px 0 0;
}

.overlay .overlay-inner .content {
    width: 100%;
    height: calc(70vh - 70px - 77px);
}

.overlay .overlay-inner .content section {

    padding: 0;
    height: 100%;
    /* height: calc(70vh - 70px - 77px); */
    padding-top: 0px;
    font-size: 20px;
    color: hsla(160, 100%, 37%, 1);
    display: flex;
}

.overlay .overlay-inner .content section input[type="text"] {
    border: solid 1px hsla(160, 100%, 37%, 1);
    background-color: white;
    height: 42px;
    padding: 0px 10px;
    font-size: 20px;
    width: 90%;
    margin-top: 1rem;
    color: hsla(160, 100%, 37%, 1);
}

.overlay .overlay-inner .content section textarea {
    border: solid 1px hsla(160, 100%, 37%, 1);
    background-color: white;
    height: 168px;
    padding: 0px 10px;
    font-size: 20px;
    width: 90%;
    margin-top: 1rem;
    color: hsla(160, 100%, 37%, 1);
}

.overlay .overlay-inner .content section div {
    width: 50%;
    padding: 1rem;
    overflow: scroll;
}

.overlay .overlay-inner .content section div.img {
    overflow: hidden;
    padding: 0;
    display: flex;
    align-items: center;
    justify-content: center;
    position: relative;
}

.overlay .overlay-inner .content section div.img img {
    height: 100%;
    /* transform: translateX(-50%); */
}

.overlay .overlay-inner .content section div.img .movie-name {
    position: absolute;
    bottom: 20px;
    left: 0;
    width: 100%;

    display: flex;
    justify-content: center;
    align-items: center;

    color: #002422;
    font-size: 30px;
    font-weight: 1500;
    font-family: 'bangers';

}



.overlay .overlay-inner .footer {
    padding: 1rem;
    border: none;
    border-top: solid 3px hsla(160, 100%, 37%, 1);
    color: white;
    font-size: 25px;
    font-weight: bold;
    display: flex;
    justify-content: space-between;
    flex-direction: row-reverse;
}

.overlay .overlay-inner .footer .btn {
    border: solid 1px hsla(160, 100%, 37%, 1);
    height: 42px;
    padding: 0px 45px;
    font-size: 20px;
}

.overlay .overlay-inner .footer .btn:disabled {
    border: solid 1px hsla(160, 100%, 37%, 0.2);
    color: hsla(160, 100%, 37%, 0.2);
}

.overlay .overlay-inner .footer .btn.btn-next {
    background-color: hsla(160, 100%, 37%, 0.2);
    color: hsla(160, 100%, 37%, 1);
}

.overlay .overlay-inner .footer .btn.btn-back {
    background-color: hsla(160, 100%, 37%, 0);
    color: hsla(160, 100%, 37%, 1);
}</style>
