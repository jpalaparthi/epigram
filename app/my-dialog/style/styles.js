import { StyleSheet } from "react-native";

const styles = StyleSheet.create({
  ViewRenderItemText: {
    minHeight: 30,
    backgroundColor: "#fff",
    color: "black",
    textAlign: "justify",
    padding: 10,
    flex: 1,
    flexWrap: "wrap",
    fontFamily: "Arial"
  },
  ViewIcon: {
    width: 40,
    alignItems: "center",
    justifyContent: "center"
  },
  RenderSeperator: {
    height: 1,
    width: "100%",
    backgroundColor: "#CED0CE",
    alignItems:"baseline"
  },
  container: {
    //flex: 1,
    //flexWrap: "wrap"
    height: "100%", // This makes container occupies full height
    //width: "100%",
}
});

export default styles;
