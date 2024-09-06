<template>
  <div class="container">
    <Navbar></Navbar>
    <main>
      <router-view />
    </main>
  </div>
</template>

<script setup>
import Navbar from "@/components/NavBar.vue";
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Noto+Sans+TC:wght@300;400&display=swap');

* {
  margin: 0;
  padding: 0;
  outline: 0;
  appearance: none;
  border: 0;
  text-decoration: none;
  list-style-type: none;
  /* box-sizing: border-box; */
  font-family: 'Noto Sans TC', sans-serif;
}

body {
  background: linear-gradient(rgba(0, 0, 0, 0.2), rgba(0, 0, 0, 0.2)), url('./assets/img/Rectangle(filter).png') no-repeat center center fixed;
  background-blend-mode: darken;
  -webkit-background-size: cover;
  -moz-background-size: cover;
  -o-background-size: cover;
  background-size: cover;
  margin: 0px;
  padding: 0px;
}

main {
  width: 85%;
  height: auto;
  padding: 1rem 0 1rem 0;
  margin: 0 auto;
}

@media screen and (max-width: 1440px) {
  main {
    width: 95%;

  }
}
</style>
