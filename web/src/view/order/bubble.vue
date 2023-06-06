<template>
  <div class="charging-animation">
    <div class="g-number">{{ chargePercentage }}%</div>
    <div class="g-container">
      <div class="g-contrast">
        <div class="g-circle"></div>
      </div>
      <div class="g-bubbles">
        <li v-for="(bubble, index) in bubbles" :key="index" class="bubble"></li>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      chargePercentage: 0,
      bubbles: Array.from({ length: 15 }), // 创建一个包含15个元素的数组
    };
  },
  mounted() {
    // 模拟充电进度增加
    setInterval(() => {
      this.chargePercentage = (this.chargePercentage + 1) % 101;
    }, 100000000);
  },
}
</script>

<style scoped>
.charging-animation {
  width: 25%;
  height: 25%;
  display: flex;
  overflow: hidden;
}

.g-number {
  position: absolute;
  width: 300px;
  top: 39%;
  text-align: center;
  font-size: 32px;
  z-index: 10;
  color: #fff;
}

.g-container {
  position: relative;
  width: 300px;
  height: 400px;
  margin: auto;
}

.g-contrast {
  filter: contrast(10) hue-rotate(0);
  width: 300px;
  height: 400px;
  overflow: hidden;
  animation: hueRotate 10s infinite linear;
}

.g-circle {
  position: relative;
  width: 300px;
  height: 300px;
  box-sizing: border-box;
  filter: blur(8px);
}

.g-circle::after {
  content: "";
  position: absolute;
  top: 40%;
  left: 50%;
  transform: translate(-50%, -50%) rotate(0);
  width: 200px;
  height: 200px;
  background-color: #00ff6f;
  border-radius: 42% 38% 62% 49%/45%;
  animation: rotate 10s infinite linear;
}

.g-circle::before {
  content: "";
  position: absolute;
  width: 176px;
  height: 176px;
  top: 40%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 50%;
  z-index: 10;
}

.g-bubbles {
  position: absolute;
  left: 50%;
  bottom: 0;
  width: 100px;
  height: 40px;
  transform: translate(-50%, 0);
  border-radius: 100px 100px 0 0;
  background-color: #00ff6f;
  filter: blur(5px);
}

li {
  position: absolute;
  border-radius: 50%;
  background: #00ff6f;
}

li:nth-child(0) {
  left: 41px;
  top: 50%;
  transform: translate(-50%, -50%);
  width: 25px;
  height: 25px;
  animation: moveToTop 7s ease-in-out -2.078s infinite;
}

li:nth-child(1) {
  left: 73px;
  top: 50%;
  transform: translate(-50%, -50%);
  width: 27px;
  height: 27px;
  animation: moveToTop 4s ease-in-out -2.749s infinite;
}
li:nth-child(2) {
  left: 30px;
  top: 50%;
  transform: translate(-50%, -50%);
  width: 22px;
  height: 22px;
  -webkit-animation: moveToTop 7s ease-in-out -3.531s infinite;
          animation: moveToTop 7s ease-in-out -3.531s infinite;
}
li:nth-child(3) {
  left: 25px;
  top: 50%;
  transform: translate(-50%, -50%);
  width: 28px;
  height: 28px;
  -webkit-animation: moveToTop 5s ease-in-out -2.831s infinite;
          animation: moveToTop 5s ease-in-out -2.831s infinite;
}
li:nth-child(4) {
  left: 30px;
  top: 50%;
  transform: translate(-50%, -50%);
  width: 26px;
  height: 26px;
  -webkit-animation: moveToTop 4s ease-in-out -4.635s infinite;
          animation: moveToTop 4s ease-in-out -4.635s infinite;
}
li:nth-child(5) {
  left: 28px;
  top: 50%;
  transform: translate(-50%, -50%);
  width: 19px;
  height: 19px;
  -webkit-animation: moveToTop 8s ease-in-out -4.057s infinite;
          animation: moveToTop 8s ease-in-out -4.057s infinite;
}
li:nth-child(6) {
  left: 20px;
  top: 50%;
  transform: translate(-50%, -50%);
  width: 19px;
  height: 19px;
  -webkit-animation: moveToTop 6s ease-in-out -2.653s infinite;
          animation: moveToTop 6s ease-in-out -2.653s infinite;
}
li:nth-child(7) {
  left: 80px;
  top: 50%;
  transform: translate(-50%, -50%);
  width: 23px;
  height: 23px;
  -webkit-animation: moveToTop 5s ease-in-out -1.79s infinite;
          animation: moveToTop 5s ease-in-out -1.79s infinite;
}
li:nth-child(8) {
  left: 76px;
  top: 50%;
  transform: translate(-50%, -50%);
  width: 26px;
  height: 26px;
  -webkit-animation: moveToTop 7s ease-in-out -0.813s infinite;
          animation: moveToTop 7s ease-in-out -0.813s infinite;
}
li:nth-child(9) {
  left: 37px;
  top: 50%;
  transform: translate(-50%, -50%);
  width: 27px;
  height: 27px;
  -webkit-animation: moveToTop 7s ease-in-out -2.532s infinite;
          animation: moveToTop 7s ease-in-out -2.532s infinite;
}
li:nth-child(10) {
  left: 24px;
  top: 50%;
  transform: translate(-50%, -50%);
  width: 20px;
  height: 20px;
  -webkit-animation: moveToTop 7s ease-in-out -4.09s infinite;
          animation: moveToTop 7s ease-in-out -4.09s infinite;
}
li:nth-child(11) {
  left: 54px;
  top: 50%;
  transform: translate(-50%, -50%);
  width: 20px;
  height: 20px;
  -webkit-animation: moveToTop 9s ease-in-out -3.521s infinite;
          animation: moveToTop 9s ease-in-out -3.521s infinite;
}

@-webkit-keyframes rotate {
  50% {
    border-radius: 45%/42% 38% 58% 49%;
  }
  100% {
    transform: translate(-50%, -50%) rotate(720deg);
  }
}
@keyframes rotate {
  50% {
    border-radius: 45%/42% 38% 58% 49%;
  }
  100% {
    transform: translate(-50%, -50%) rotate(720deg);
  }
}
@-webkit-keyframes moveToTop {
  90% {
    opacity: 1;
  }
  100% {
    opacity: 0.1;
    transform: translate(-50%, -180px);
  }
}
@keyframes moveToTop {
  90% {
    opacity: 1;
  }
  100% {
    opacity: 0.1;
    transform: translate(-50%, -180px);
  }
}
@-webkit-keyframes hueRotate {
  100% {
    filter: contrast(15) hue-rotate(360deg);
  }
}
@keyframes hueRotate {
  100% {
    filter: contrast(15) hue-rotate(360deg);
  }
}
</style>

