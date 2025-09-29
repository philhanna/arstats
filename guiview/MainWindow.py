from PySide6.QtWidgets import QApplication, QMainWindow
from GuiView import Ui_MainWindow  # This matches the <class> in your .ui file
import subprocess
import sys

class MainWindow(QMainWindow):
    def __init__(self):
        super().__init__()
        self.ui = Ui_MainWindow()
        self.ui.setupUi(self)

        # Example: set dynamic text in a label
        result = subprocess.getoutput(["arstats"])
        lines = result.splitlines()
        self.ui.listWidget.clear()
        self.ui.listWidget.addItems(lines)

if __name__ == "__main__":
    app = QApplication(sys.argv)
    window = MainWindow()
    window.show()
    sys.exit(app.exec())
