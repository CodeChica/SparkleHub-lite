document.addEventListener('DOMContentLoaded', (event) => {
  window.app = new Vue({
    el: '#app',
    data: {
      intervalId: null,
      isSending: false,
      errorMessage: "",
      sparkle: "",
      sparkles: [],
      searchTerm: ""
    },
    created: function() {
      this.reload();
      this.intervalId = setInterval(() => this.reload(), 30000);
    },
    destroyed: function() {
      if (this.intervalId)
        clearInterval(this.intervalId);
      this.intervalId = null;
    },
    computed: {
      heading: function() {
        return this.sparkles.length == 0 ? "No Sparkles Sent" : "Recent Sparkles";
      },
      recentSparkles: function() {
        return this.sparkles.reverse();
      },
      isDisabled: function() {
        return this.isSending || !this.isValid();
      },
      filteredSparkles: function() {
        if (this.searchTerm === ""){
          return this.sparkles.reverse();
        }
        return this.recentSparkles.filter((sparkle) => {
          return sparkle.sparklee.includes(this.searchTerm)
            || sparkle.reason.includes(this.searchTerm);
        });
      }
    },
    watch: {
      sparkle: function() {
        this.errorMessage = "";
      },
    },
    methods: {
      reload: function() {
        fetch("/v2/sparkles")
          .then((response) => response.json())
          .then((json) => this.sparkles = json.data.map(x => x.attributes))
          .catch((json) => console.error(json.errors));
      },
      isValid: function() {
        return this.sparkle.length > 0;
      },
      startConfetti: function() {
        let message = document.querySelector('#sparkle-sent-message');
        message.classList.remove("hidden");
        message.start();

        let container = document.querySelector('.confetti-container');

        for(let index = 255; index >= 0; index--) {
          let div = document.createElement("div");
          div.classList.add("confetti-" + index.toString())
          container.appendChild(div);
        }

        setTimeout(() => this.removeConfetti(), 12000);
      },
      removeConfetti: function() {
        let element = document.querySelector('.confetti-container')
        let message = document.querySelector('#sparkle-sent-message');
        message.classList.add("hidden");

        while (element.firstChild) {
          element.removeChild(element.firstChild);
        }
      },
      submitSparkle: function() {
        this.isSending = true;
        fetch("/sparkles.json", {
          method: "POST",
          mode: "cors",
          cache: "no-cache",
          headers: { "Content-Type": "application/json" },
          redirect: "follow",
          body: JSON.stringify({ body: this.sparkle })
        }).then((response) => {
          response.json().then((json) => {
            this.isSending = false;
            if (response.ok) {
              this.startConfetti();
              this.sparkles.push(json);
              this.sparkle = "";
            } else {
              this.errorMessage = json["error"];
            }
          })
        }).catch((error) => console.error(error));
      }
    }
  })
});
// Simple example, see optional options for more configuration.
const pickr = Pickr.create({
  el: '.color-picker',
  theme: 'classic', // or 'monolith', or 'nano'

  swatches: [
      'rgba(244, 67, 54, 1)',
      'rgba(233, 30, 99, 0.95)',
      'rgba(156, 39, 176, 0.9)',
      'rgba(103, 58, 183, 0.85)',
      'rgba(63, 81, 181, 0.8)',
      'rgba(33, 150, 243, 0.75)',
      'rgba(3, 169, 244, 0.7)',
      'rgba(0, 188, 212, 0.7)',
      'rgba(0, 150, 136, 0.75)',
      'rgba(76, 175, 80, 0.8)',
      'rgba(139, 195, 74, 0.85)',
      'rgba(205, 220, 57, 0.9)',
      'rgba(255, 235, 59, 0.95)',
      'rgba(255, 193, 7, 1)'
  ],

  components: {

      // Main components
      preview: true,
      opacity: true,
      hue: true,

      // Input / output Options
      interaction: {
          hex: true,
          rgba: true,
          hsla: true,
          hsva: true,
          cmyk: true,
          input: true,
          clear: true,
          save: true
      }
  }
});

let uploadButton = document.getElementById ("upload-button");   
let chosenImage = document.getElementbyId   ("chosen-image");   
let fileName = document.getElementbyId   ("file-name");   

uploadButton.onchange = () => {  let reader = newFileReader();   
  reader.ReadAsDatatURL (uploadButton.files[0]);   
  read.onload = () => {   
    chosenIimage.setAttribute('src", reader. result); }   
    fileName.textContent = uploadButton.files[0]. 
    name;