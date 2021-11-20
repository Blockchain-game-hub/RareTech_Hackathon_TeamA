<template>
  <div>
    <div class="relative">
      <main>
        <div class="flex justify-center ml-32">
          <canvas
            class="bg-white nes-container is-rounded"
            id="gameCanvas"
            ref="gameCanvas"
            @mousemove="draw"
            @mousedown="dragStart"
            @mouseup="dragEnd"
            @mouseout="dragEnd"
          ></canvas>
        </div>
      </main>
      <div class="absolute top-0 ml-10">
        <i class="snes-jp-logo -mt-2 ml-10"></i>
        <div
          class="
            flex flex-col
            w-32
            nes-container
            is-rounded
            bg-white
            items-center
          "
        >
          <button
            class="nes-btn mt-3"
            @click="penMode"
            :class="{ modeOn: isPenMode }"
          >
            <i class="gg-pen mb-2 ml-1 mr-1"></i>
          </button>
          <button
            class="nes-btn mt-3"
            @click="eraserMode"
            :class="{ modeOn: isPenMode == false }"
          >
            <i class="gg-erase -ml-1"></i>
          </button>
          <input
            class="mt-3"
            type="color"
            id="color"
            v-model="lineStatus.color"
          />
          <div class="mt-3 ml-4">
            <input
              class="press-start-2"
              type="number"
              min="1"
              max="10"
              id="line-width"
              v-model="lineStatus.lineWidth"
            />
          </div>

          <button class="mt-3 p-2 nes-btn" @click="dargRedo">
            <i class="gg-redo"></i>
          </button>
          <button class="mt-3 p2 nes-btn" @click="dargUndo">
            <i class="gg-undo"></i>
          </button>
          <button class="mt-3 nes-btn is-error press-start-2" @click="allClear">
            clear
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {
  defineComponent,
  onMounted,
  reactive,
  ref,
} from "@nuxtjs/composition-api";
import io from 'socket.io-client'

interface positionStack {
  redoDataStack: any[];
  undoDataStack: any[];
}

export default defineComponent({
  props: {
    io: {
      type: io,
      required: true
    }
  },
  data() {
    return {
      socket: ''
    }
  },
  setup(socket) {
    //   線の情報、初期値
    const lineStatus = reactive({
      lineWidth: 5,
      color: "black",
    });
    const canvasDefultStatus = {
      width: 900,
      height: 400,
    };

    // canvasの情報,初期値null
    const context = ref<CanvasRenderingContext2D | any>();
    onMounted(() => {
      //   canvasの初期値を代入
      const canvas = <HTMLCanvasElement>document.getElementById("gameCanvas");
      context.value = canvas.getContext("2d");
      canvas.width = canvasDefultStatus.width;
      canvas.height = canvasDefultStatus.height;
      context.value.lineCap = "round";
      //    変更可能なcanvasの値
      context.value.lineWidth = lineStatus.lineWidth;
    });

    // 1.0 描画
    // 1.1.書いているかいないか判定
    let isDrag: Boolean = false;

    // 1.2絵を描く関数(描画中)
    const draw = (e: any) => {
      let x = e.layerX - 23;
      let y = e.layerY - 20;
      if (!isDrag) return;
      context.value.lineTo(x, y);
      context.value.stroke();
    };

    // 1.3 描画開始
    const dragStart = (e: any) => {
      let x = e.layerX - 23;
      let y = e.layerY - 20;
      isDrag = true;
      //   3.1 元に戻すボタン用にimageDataを保存しておく
      positionStack.undoDataStack.push(
        context.value.getImageData(
          0,
          0,
          canvasDefultStatus.width,
          canvasDefultStatus.height
        )
      );
      context.value.beginPath();
      context.value.lineTo(x, y);
      context.value.stroke();
      context.value.getImageData(
        0,
        0,
        canvasDefultStatus.width,
        canvasDefultStatus.height
      );
      //   3.5戻る機能 制限 30回まで 以降削除
      if (positionStack.undoDataStack.length >= 30) {
        positionStack.undoDataStack.pop();
      }
      //   3.10 進む用の配列を空にしておく
      positionStack.redoDataStack = [];

      //   色変更
      //   モード変更
      if (isPenMode.value) {
        context.value.strokeStyle = lineStatus.color;
      } else {
        context.value.strokeStyle = "#FFFFFF";
      }
      //   太さ変更
      context.value.lineWidth = lineStatus.lineWidth;
    };

    // 1.4 描画終了(mouseup,mouseout)
    const dragEnd = () => {
      context.value.closePath();
      isDrag = false;
    };

    // 2.0モード変更 pen or 消しゴム
    const isPenMode = ref<boolean>(true);
    // 2.1 penモードに変更
    const penMode = () => {
      isPenMode.value = true;
    };
    // 2.2 消しゴムモードに変更
    const eraserMode = () => {
      isPenMode.value = false;
    };

    // 3.0 1コ進む,1コ戻る機能
    const positionStack = <positionStack>reactive({
      redoDataStack: [],
      undoDataStack: [],
    });

    // 3.2 1コ戻る機能
    const dargUndo = () => {
      if (positionStack.undoDataStack.length > 0) {
        // 3.3 進む用にデータをスタックしておく
        positionStack.redoDataStack.push(
          context.value.getImageData(
            0,
            0,
            canvasDefultStatus.width,
            canvasDefultStatus.height
          )
        );
        //   3.4undoDataStackからイメージデータを取得 描画
        let imageData = positionStack.undoDataStack.pop();
        context.value.putImageData(imageData, 0, 0);
      }
    };

    // 3.6進む機能
    const dargRedo = () => {
      if (positionStack.redoDataStack.length > 0) {
        // 3.7 戻る用にデータをスタックしておく
        positionStack.undoDataStack.push(
          context.value.getImageData(
            0,
            0,
            canvasDefultStatus.width,
            canvasDefultStatus.height
          )
        );
        //   3.8redoDataStackからイメージデータを取得 描画
        let imageData = positionStack.redoDataStack.pop();
        context.value.putImageData(imageData, 0, 0);
      }
    };

    // 4.0全て消す
    const allClear = () => {
      const verification = confirm("本当に全て消しますか？");
      if (verification) {
        context.value.clearRect(
          0,
          0,
          canvasDefultStatus.width,
          canvasDefultStatus.height
        );
      }
    };

    return {
      draw,
      dragStart,
      dragEnd,
      lineStatus,
      penMode,
      eraserMode,
      dargUndo,
      dargRedo,
      allClear,
      isPenMode,
    };
  },
});
</script>

<style scoped>
.modeOn {
  background-color: rgba(141, 134, 134, 0.555);
}
</style>
