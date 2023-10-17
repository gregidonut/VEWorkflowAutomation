import IndexPage from "./indexPage.js";
import EditPage from "./editPage.js";

type PageRegistry = {
    [route: string]: () => void;
};
const pageRegistry: PageRegistry = { "/": IndexPage, "/edit": EditPage };

for (const [route, fn] of Object.entries(pageRegistry)) {
    if (window.location.pathname == route) {
        fn();
    }
}
