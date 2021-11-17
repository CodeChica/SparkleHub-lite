document.addEventListener('DOMContentLoaded', (event) => {
  window.app = new Vue({
    el: '#app',
    data: {
      intervalId: null,
      isSending: false,
      errorMessage: "",
      sparkle: "",
      sparkles: [],
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
        return this.sparkles.length == 0 ? "No Sparkles Yet" : "Recent Sparkles";
      },
      recentSparkles: function() {
        return this.sparkles.reverse();
      },
      isDisabled: function() {
        return this.isSending || !this.isValid();
      },
    },
    watch: {
      sparkle: function() {
        this.errorMessage = "";
      },
    },
    methods: {
      reload: function() {
        fetch("/sparkles.json")
          .then((response) => response.json())
          .then((data) => this.sparkles = data)
          .catch((error) => console.error(error));
      },
      isValid: function() {
        return this.sparkle.length > 0;
      },
      sparkleParty: function() {
        let container = document.querySelector('.confetti-container')

        for(let index = 255; index >= 0; index--) {
          let div = document.createElement("div");
          div.classList.add("confetti-" + index.toString())
          container.appendChild(div);
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
              this.sparkleParty();
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
