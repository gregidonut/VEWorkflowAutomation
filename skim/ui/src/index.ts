import IndexPage from "/static/dist/indexPage.js";
import EditPage from "/static/dist/editPage.js";

if (window.location.pathname == "/") {
    IndexPage();
}

if (window.location.pathname == "/edit") {
    EditPage();
}
