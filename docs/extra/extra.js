const backgrounds = [
  "https://user-images.githubusercontent.com/3319450/209234376-3c4f08a5-80a9-4b38-ad40-a4aa9b6a1764.png",
  "https://user-images.githubusercontent.com/3319450/209234411-25086f15-d076-4e47-bc96-0c72518fbf87.png",
  "https://user-images.githubusercontent.com/3319450/209234500-d1a0c432-c535-4261-b919-0a5471965aef.png",
  "https://user-images.githubusercontent.com/3319450/209234471-48435c70-3047-4260-8220-dc9e912377a7.png",
  "/assets/kelethin-render-min.png",
  "https://raw.githubusercontent.com/Akkadius/eq-asset-preview/master/assets/wallpaper/faydark.png",
  "https://raw.githubusercontent.com/Akkadius/eq-asset-preview/master/assets/wallpaper/soldungb.png",
  "https://raw.githubusercontent.com/Akkadius/eq-asset-preview/master/assets/wallpaper/sebilis.png",
  "https://raw.githubusercontent.com/Akkadius/eq-asset-preview/master/assets/wallpaper/rivervale.png",
  "https://raw.githubusercontent.com/Akkadius/eq-asset-preview/master/assets/wallpaper/qeynos.png",
  "https://raw.githubusercontent.com/Akkadius/eq-asset-preview/master/assets/wallpaper/oasis.png",
  "https://raw.githubusercontent.com/Akkadius/eq-asset-preview/master/assets/wallpaper/lavastorm.png",
  "https://raw.githubusercontent.com/Akkadius/eq-asset-preview/master/assets/wallpaper/freeport.png",
]

for (let bg of backgrounds) {
  // console.log("loading " + bg)
  let img = new Image();
  img.src = bg;
}

function shuffleBackgrounds() {
  let index = 1;
  for (let bg of shuffle(backgrounds)) {
    document.body.style.setProperty("--bg" + index, "url(" + bg + ")");
    // console.log("setting bg" + index + " to " + bg)
    index++;
  }
}

function shuffle(array) {
  let currentIndex = array.length, randomIndex;

  // While there remain elements to shuffle.
  while (currentIndex !== 0) {

    // Pick a remaining element.
    randomIndex = Math.floor(Math.random() * currentIndex);
    currentIndex--;

    // And swap it with the current element.
    [array[currentIndex], array[randomIndex]] = [
      array[randomIndex], array[currentIndex]];
  }

  return array;
}

shuffleBackgrounds()

window.onload = function () {
  onLoad()
};

let origOpen                  = XMLHttpRequest.prototype.open;
XMLHttpRequest.prototype.open = function () {
  this.addEventListener('loadend', function () {
    onLoad()
  });
  origOpen.apply(this, arguments);
};

function onLoad() {
  // shuffleBackgrounds()
}
