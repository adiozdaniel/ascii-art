package renders

var Tmpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
	<style>
	html,
body {
	background-color: #333;
	font-family: Arial, sans-serif;
	text-align: center;
	margin: 0;
	padding: 0;
	color: bisque;
	font-size: 1.2em;
	height: 100vh;
}

h1,
h2 {
	color: bisque;
}

h1 {
	transform: scale(90%);
}

.btn {
	height: fit-content;
	position: relative;
	text-decoration: none;
	color: bisque;
	letter-spacing: 0.2em;
	cursor: pointer;
	background-color: #235525;
	margin: 2% auto;
	padding: 2% 5%;
	width: 30%;
	border: none;
	border-radius: 4px;
	box-shadow: 0 3px 10px rgb(1, 2, 19);
}

.btn:hover {
	background-color: #5b7e5c;
	color: rgb(1, 2, 19);
	text-decoration: none;
	box-shadow: 0 2px 15px rgb(1, 2, 19);
}

.navbar {
	background-color: #222;
	overflow: hidden;
	padding: 0;
	margin: 0;
	box-shadow: 0 8px 6px rgb(1, 2, 19);
	display: flex;
	font-size: 0.8em;
	justify-content: space-between;
	align-items: center;
	transition: transform 0.3s ease;
}

/* Hamburger Icon Styles */
.menu-toggle {
	display: none;
}

.hamburger {
	display: none;
	flex-direction: column;
	cursor: pointer;
	padding: 10px;
}

.hamburger .bar {
	width: 30px;
	height: 3px;
	background-color: bisque;
	margin: 5px 0;
	transition: 0.3s;
}

/* User Icon Styles */
.user-icon {
	transition: 0.3s;
	font-size: 0.8em;
	display: flex;
	align-items: center;
	padding: 0.6em;
	transition: transform 0.3s ease;
}

.user-icon img {
	display: block;
	width: 40px;
	height: 40px;
	cursor: pointer;
	transition: 0.3s;
	border-radius: 50%;
	transition: transform 0.3s ease;
}

.username {
	font-weight: bold;
	color: #4c9850;
	transition: 0.3s ease;
}

.navbar ul {
	list-style: none;
	margin: 0;
  	padding: 0;
	width: 55%;
	display: flex;
	align-items: center;
	justify-content: space-evenly;
}

.navbar ul li {
	width: 20%;
	margin: 0.5%;
	transition: transform 0.3s ease;
}

.navbar ul li:hover {
	background-color: #4c9850;
}

.navbar ul li a {
	display: block;
  	text-align: center;
  	padding: 14px 5px;
	width: 100%;
	color: bisque;
	text-decoration: none;
	transition: 0.3s;
}

.navbar ul li a:hover {
	color: rgb(1, 2, 19);
}

.navbar #logout {
	letter-spacing: 0.2em;
	font-size: 0.7em;
	color: bisque;
	background-color: rgb(139, 3, 166);
	border-radius: 0.6em;
	padding: 0.2em 0.9em;
	transition: transform 0.3s ease;
}

.navbar #logout:hover {
	background-color: rgb(76, 2, 91);
}

.user-icon.shrink {
	transform: scale(0.6);
}

/* Show hamburger icon on small screens */
@media (max-width: 544px) {
	h1 {
		transform: scale(75%);
	}
	.hamburger {
		display: flex;
	}

	.navbar ul {
		display: none;
		flex-direction: column;
		position: absolute;
		top: 60px;
		right: 0;
		width: 100%;
	}

	.menu-toggle:checked + .hamburger #bar1 {
		transform: rotate(45deg);
		position: relative;
		top: 15px;
		transition: transform 0.4s ease;
	}

	.menu-toggle:checked + .hamburger #bar2 {
		opacity: 0;
		transition: transform 0.4s ease;
	}

	.menu-toggle:checked + .hamburger #bar3 {
		transform: rotate(-45deg);
		position: relative;
		top: -10px;
		transition: transform 0.4s ease;
	}

	.menu-toggle:checked ~ #nav-menu {
		display: flex;
		z-index: 500;
		opacity: 0.9;
		background-color: #222;
		transition: transform 0.4s ease;
		transition: background-color 1s ease;
	}

	.navbar ul li {
		padding: 10px;
		width: 50%;
		background-color: rgb(19, 9, 1);
		border-radius: 5%;
		transition: background-color 0.5s ease;
	}
}

@media (max-width: 280px) {
	.navbar ul li {
		width: 60%;
	}
}

.wrapper {
	bottom: 0%;
	margin-top: 10%;
	width: 100%;
	position: relative;
}

svg {
	padding: 10px;
}

.authors {
	margin: 2em auto -8em auto;
}

.authors p {
	margin-top: -1.5em;
	font-size: 0.7em;
	color: bisque;
	text-shadow: 0 0 0.5em rgb(254, 228, 196);
	text-shadow: 1px 1px 2px rgb(1, 2, 19), 0 0 5px #222;
}

.author {
	display: inline-block;
	margin: 5%;
	text-align: center;
	font-size: medium;
}

.author img {
	border-radius: 50%;
	width: 80px;
	height: 80px;
	object-fit: cover;
	box-shadow: 2px 8px 15px rgb(1, 2, 19);
}

.author-name {
	margin-top: 10px;
	color: bisque;
	font-size: 1.2em;
}

.my-footer {
	background-color: #222;
	padding: 1% 20%;
	margin: 0 auto;
	text-align: center;
	color: bisque;
	display: flex;
	flex-wrap: wrap;
	align-items: center;
	justify-content: space-between;
	align-items: center;
	border-top: 1px solid #444;
	border-bottom: 1px solid #444;
}

.my-footer svg {
	fill: bisque;
	margin: 0 10px;
}

.my-footer span {
	font-size: 0.8em;
}

@media screen and (max-width: 544px) {
	p {
		padding: 0;
		font-size: 0.9em;
	}

	h2 {
		font-size: 1.2em;
	}

	.authors p {
		margin-top: -1.2em;
		font-size: 0.5em;
	}
}

	</style>

    <title>Server Error</title>
</head>
<body>
    <div class="navbar">
        <!-- Hamburger Icon -->
        <input
            type="checkbox"
            id="menu-toggle"
            class="menu-toggle"
        />
        <label
            for="menu-toggle"
            class="hamburger"
        >
            <div
                id="bar1"
                class="bar"
            ></div>
            <div
                id="bar2"
                class="bar"
            ></div>
            <div
                id="bar3"
                class="bar"
            ></div>
        </label>
    
        <!-- Navigation Menu -->
        <ul id="nav-menu">
            <li id="home">
                <a href="/">Home</a>
            </li>
            <li id="ascii">
                <a href="/ascii-art">Ascii-Art</a>
            </li>
            <li id="about">
                <a href="/about">About</a>
            </li>
            <li id="contact">
                <a href="/contact">Contact</a>
            </li>
        </ul>

        <!-- server error content -->
    <div>
        <h1>500 Oops! Server Error 🙁</h1>
			<h2>Something went wrong.</h2>
			<h3>{{.Error}}</h3>
			<a href="/" title="Go back to the home page" class="btn">
			<h1>Home</h1>
			</a>
    </div>
        <!-- content end -->

    <div class="wrapper">
        <!-- Authors Section -->
        <div class="authors">
            <h3>Authors</h3>
            <div class="author">
                <img
                    src="https://avatars.githubusercontent.com/u/77047643?v=4 "
                    alt="Author 1"
                />
                <div class="author-name">Josephine Opondo</div>
            </div>
            <div class="author">
                <img
                    src="https://lh3.googleusercontent.com/a/ACg8ocLUKAW3QwBqLDqDcmkFTC3wmCPq0dd25wVFn3CPEkCfhQQme9Lx=s288-c-no"
                    alt="Author 2"
                />
                <div class="author-name">Andrew Osindo</div>
            </div>
            <div class="author">
                <img
                    src="https://avatars.githubusercontent.com/u/42722945?v=4"
                    alt="Author 3"
                />
                <div class="author-name">Adioz Eshitemi</div>
            </div>
            <p>
                Stay updated with the latest news, features, and fun stuff. Connect with
                us on:
            </p>
        </div>
        <!-- Authors-End -->
    
        <!-- Footer-Start -->
        <svg
            xmlns="http://www.w3.org/2000/svg"
            class="d-none"
        >
            <symbol
                id="facebook"
                viewBox="0 0 16 16"
            >
                <path
                    d="M16 8.049c0-4.446-3.582-8.05-8-8.05C3.58 0-.002 3.603-.002 8.05c0 4.017 2.926 7.347 6.75 7.951v-5.625h-2.03V8.05H6.75V6.275c0-2.017 1.195-3.131 3.022-3.131.876 0 1.791.157 1.791.157v1.98h-1.009c-.993 0-1.303.621-1.303 1.258v1.51h2.218l-.354 2.326H9.25V16c3.824-.604 6.75-3.934 6.75-7.951z"
                />
            </symbol>
            <symbol
                id="instagram"
                viewBox="0 0 16 16"
            >
                <path
                    d="M8 0C5.829 0 5.556.01 4.703.048 3.85.088 3.269.222 2.76.42a3.917 3.917 0 0 0-1.417.923A3.927 3.927 0 0 0 .42 2.76C.222 3.268.087 3.85.048 4.7.01 5.555 0 5.827 0 8.001c0 2.172.01 2.444.048 3.297.04.852.174 1.433.372 1.942.205.526.478.972.923 1.417.444.445.89.719 1.416.923.51.198 1.09.333 1.942.372C5.555 15.99 5.827 16 8 16s2.444-.01 3.298-.048c.851-.04 1.434-.174 1.943-.372a3.916 3.916 0 0 0 1.416-.923c.445-.445.718-.891.923-1.417.197-.509.332-1.09.372-1.942C15.99 10.445 16 10.173 16 8s-.01-2.445-.048-3.299c-.04-.851-.175-1.433-.372-1.941a3.926 3.926 0 0 0-.923-1.417A3.911 3.911 0 0 0 13.24.42c-.51-.198-1.092-.333-1.943-.372C10.443.01 10.172 0 7.998 0h.003zm-.717 1.442h.718c2.136 0 2.389.007 3.232.046.78.035 1.204.166 1.486.275.373.145.64.319.92.599.28.28.453.546.598.92.11.281.24.705.275 1.485.039.843.047 1.096.047 3.231s-.008 2.389-.047 3.232c-.035.78-.166 1.203-.275 1.485a2.47 2.47 0 0 1-.599.919c-.28.28-.546.453-.92.598-.28.11-.704.24-1.485.276-.843.038-1.096.047-3.232.047s-2.39-.009-3.233-.047c-.78-.036-1.203-.166-1.485-.276a2.478 2.478 0 0 1-.92-.598 2.48 2.48 0 0 1-.6-.92c-.109-.281-.24-.705-.275-1.485-.038-.843-.046-1.096-.046-3.233 0-2.136.008-2.388.046-3.231.036-.78.166-1.204.276-1.486.145-.373.319-.64.599-.92.28-.28.546-.453.92-.598.282-.11.705-.24 1.485-.276.738-.034 1.024-.044 2.515-.045v.002zm4.988 1.328a.96.96 0 1 0 0 1.92.96.96 0 0 0 0-1.92zm-4.27 1.122a4.109 4.109 0 1 0 0 8.217 4.109 4.109 0 0 0 0-8.217zm0 1.441a2.667 2.667 0 1 1 0 5.334 2.667 2.667 0 0 1 0-5.334z"
                />
            </symbol>
            <symbol
                id="twitter"
                viewBox="0 0 16 16"
            >
                <path
                    d="M5.026 15c6.038 0 9.341-5.003 9.341-9.334 0-.14 0-.282-.006-.422A6.685 6.685 0 0 0 16 3.542a6.658 6.658 0 0 1-1.889.518 3.301 3.301 0 0 0 1.447-1.817 6.533 6.533 0 0 1-2.087.793A3.286 3.286 0 0 0 7.875 6.03a9.325 9.325 0 0 1-6.767-3.429 3.289 3.289 0 0 0 1.018 4.382A3.323 3.323 0 0 1 .64 6.575v.045a3.288 3.288 0 0 0 2.632 3.218 3.203 3.203 0 0 1-.865.115 3.23 3.23 0 0 1-.614-.057 3.283 3.283 0 0 0 3.067 2.277A6.588 6.588 0 0 1 .78 13.58a6.32 6.32 0 0 1-.78-.045A9.344 9.344 0 0 0 5.026 15z"
                />
            </symbol>
        </svg>
    
        <div class="b-example-divider"></div>
    
        <footer class="my-footer">
            <div class="">
                <span>&copy; Zone01 Kisumu</span>
            </div>
            <div>
                <svg
                    class="bi"
                    width="24"
                    height="24"
                >
                    <a
                        href="https://twitter.com/Zone01Kisumu"
                        target="_blank"
                        rel="noopener noreferrer"
                    >
                        <use xlink:href="#twitter" />
                    </a>
                </svg>
            </div>
            <div>
                <svg
                    class="bi"
                    width="24"
                    height="24"
                >
                    <a
                        href="https://www.instagram.com/zone01_kisumu/"
                        target="_blank"
                        rel="noopener noreferrer"
                    >
                        <use xlink:href="#instagram" />
                    </a>
                </svg>
            </div>
            <div>
                <svg
                    class="bi"
                    width="24"
                    height="24"
                >
                    <a
                        href="https://web.facebook.com/Zone01Kisumu"
                        target="_blank"
                        rel="noopener noreferrer"
                    >
                        <use xlink:href="#facebook" />
                    </a>
                </svg>
            </div>
        </footer>
        <!-- Footer-End -->
    </div>
</body>
</html>`
