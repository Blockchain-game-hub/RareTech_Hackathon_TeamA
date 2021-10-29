<!-- Please remove this file from your project -->
<template>
  <div>
    <div>
      <h1 class="title text-center text-5xl mt-10">GAMETITLE</h1>
    </div>
    <div
      class="min-h-screen w-screen flex flex-col justify-center items-center"
    >
      <div
        class="
          bg-white
          shadow-none
          sm:shadow-lg
          px-8
          sm:px-12
          w-full
          xs:w-full
          sm:w-8/12
          md:w-7/12
          lg:w-7/12
          xl:w-2/6
          h-screen
          sm:h-auto
          py-8
          rounded-lg
          mb-20
        "
      >
        <div class="text-center w-full font-bold text-3xl text-gray-600 p-4">
          LOGIN
        </div>
        <div class="w-full bg-gray-200 my-3" style="height: 1px"></div>
        <form>
          <div class="flex flex-col gap-4 px-0 py-4">
            <div>
              <label class="text-gray-700">Email address</label>
              <!-- <AtSymbolIcon class="font-medium text-2xl text-gray-600 absolute p-2.5 px-3 w-11" /> -->
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="
                  font-medium
                  text-2xl text-gray-600
                  absolute
                  p-2.5
                  px-3
                  w-11
                "
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M16 12a4 4 0 10-8 0 4 4 0 008 0zm0 0v1.5a2.5 2.5 0 005 0V12a9 9 0 10-9 9m4.5-1.206a8.959 8.959 0 01-4.5 1.207"
                />
              </svg>
              <input
                class="py-2 pl-10 border border-gray-200 w-full"
                placeholder="Email address"
                type="email"
              />
            </div>
            <div>
              <label class="text-gray-700">Password</label>
              <!-- <LockClosedIcon class="font-medium text-2xl text-gray-600 absolute p-2.5 px-3 w-11" /> -->
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="
                  font-medium
                  text-2xl text-gray-600
                  absolute
                  p-2.5
                  px-3
                  w-11
                "
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"
                />
              </svg>
              <input
                class="py-2 pl-10 border border-gray-200 w-full"
                placeholder="Password"
                type="password"
              />
            </div>
            <div class="w-full flex flex-row gap-2">
              <button
                class="
                  border border-indigo-500
                  hover:bg-indigo-500 hover:text-white
                  duration-100
                  ease-in-out
                  w-6/12
                  text-indigo-500
                  p-0
                  flex flex-row
                  justify-center
                  items-center
                  gap-1
                "
                type="submit"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-5 w-5"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M11 16l-4-4m0 0l4-4m-4 4h14m-5 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h7a3 3 0 013 3v1"
                  />
                </svg>
                Login
              </button>
              <button
                class="
                  border border-indigo-500
                  hover:bg-indigo-500 hover:text-white
                  duration-100
                  ease-in-out
                  w-6/12
                  text-indigo-500
                  p-2
                  flex flex-row
                  justify-center
                  items-center
                  gap-1
                "
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="w-5"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"
                  />
                </svg>
                Sign-up
              </button>
            </div>
            <div class="my-2 flex flex-row justify-center">
              <span class="absolute bg-white px-4">or</span>
              <div class="w-full bg-gray-200 mt-3" style="height: 1px"></div>
            </div>
            <div class="w-full flex flex-col gap-2">
              <button
                class="
                  bg-red-500
                  text-white
                  w-full
                  p-2
                  flex flex-row
                  justify-center
                  gap-2
                  items-center
                  hover:bg-red-600
                  duration-100
                  ease-in-out
                "
              >
                Sign-in with Google
              </button>
              <button
                class="
                  bg-blue-400
                  text-white
                  w-full
                  p-2
                  flex flex-row
                  justify-center
                  gap-2
                  items-center
                  hover:bg-blue-500
                  duration-100
                  ease-in-out
                "
              >
                Sign-in with Twitter
              </button>
            </div>
            <div class="w-full flex flex-row justify-end">
              <a href="#">Forgot password</a>
            </div>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>
<script>
import "@fontsource/bungee-shade";
export default {};
</script>
<style>
.title {
  font-family: "Bungee Shade";
}
</style>
