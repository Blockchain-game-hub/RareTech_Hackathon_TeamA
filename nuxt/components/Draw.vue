<template>
  <div>
    <main>
      <div>
        <canvas
          class="bg-white"
          id="gameCanvas"
          ref="gameCanvas"
          @mousemove="draw"
          @mousedown="dragStart"
          @mouseup="dragEnd"
          @mouseout="dragEnd"
        ></canvas>
      </div>
    </main>
    <div>
      <input type="color" id="color" v-model="lineStatus.color" />
      <input
        type="number"
        min="1"
        max="10"
        id="line-width"
        v-model="lineStatus.lineWidth"
      />
      <button class="nes-btn" @click="penMode">pen</button>
      <button class="nes-btn" @click="eraserMode">eraser</button>
      <button @click="dargRedo">進む</button>
      <button @click="dargUndo">戻る</button>
      <button @click="allClear">全消し</button>
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

interface positionStack {
  redoDataStack: any[];
  undoDataStack: any[];
}

export default defineComponent({
  setup() {
    //   線の情報、初期値
    const lineStatus = reactive({
      lineWidth: 5,
      color: "black",
    });
    const canvasDefultStatus = {
      width: 1000,
      height: 400,
    };

    // canvasの情報,初期値null
    const context = ref<CanvasRenderingContext2D>();

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
      let x = e.layerX;
      let y = e.layerY;
      if (!isDrag) return;
      context.value.lineTo(x, y);
      context.value.stroke();
    };

    // 1.3 描画開始
    const dragStart = (e) => {
      let x = e.layerX;
      let y = e.layerY;
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
      //   3.5戻る機能 制限 10回まで 以降削除
      if (positionStack.undoDataStack.length >= 10) {
        positionStack.undoDataStack.pop();
      }
      //   3.10 進む用の配列を空にしておく
      positionStack.redoDataStack = [];

      //   色変更
      //   モード変更
      if (canvasMode === "pen") {
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
    let canvasMode: string = "pen";
    // 2.1 penモードに変更
    const penMode = () => {
      canvasMode = "pen";
    };
    // 2.2 消しゴムモードに変更
    const eraserMode = () => {
      canvasMode = "eraser";
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
      const confilm = confirm("本当に全て消しますか？");
      if (confirm) {
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
    };
  },
});
</script>
