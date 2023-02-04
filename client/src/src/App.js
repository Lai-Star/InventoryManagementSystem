import "./App.css";

const styles = {
  container: {
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    justifyContent: "center",
    height: "100vh",
    background:
      "linear-gradient(to right, #e84c3d, #f73859, #f73859, #f73859, #f73859, #e84c3d)",
  },
  header: {
    fontSize: 32,
    fontWeight: "bold",
    color: "#fff",
  },
  text: {
    fontSize: 16,
    color: "#fff",
    marginTop: 20,
  },
};

function App() {
  return (
    <div style={styles.container}>
      <h1 style={styles.header}>
        Welcome to Inventory Management System Home Page
      </h1>
      <p style={styles.text}>This is an application testing home page.</p>
    </div>
  );
}

export default App;
