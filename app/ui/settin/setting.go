package settin

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func UI() *widgets.QMainWindow {

	window := widgets.NewQMainWindow(nil, 0)
	// 窗口标题
	window.SetWindowTitle("团子翻译器 Ver4.0 - 设置")
	// 窗口尺寸, 最大最小尺寸
	window.Resize2(404, 576)

	// 窗口样式
	window.SetStyleSheet(`QWidget { font: 10pt \"华康方圆体W7\";
                                               background-image: url(config/static/背景.jpg);
	                                           background-repeat: no-repeat;
                                               background-size: cover;}`)

	// 窗口图标
	window.SetWindowIcon(gui.NewQIcon5("config/static/图标.ico"))
	// 鼠标样式
	pixMap := gui.NewQPixmap3("config/static/光标.png", "", core.Qt__AutoColor)
	cursor := gui.NewQCursor4(pixMap, 0, 0)
	window.SetCursor(cursor)

	// 布局窗口组件载体
	layout := widgets.NewQHBoxLayout()

	// 顶部工具栏
	tabWidget := widgets.NewQTabWidget(window)
	tabWidget.SetGeometry2(-2, 0, 410, 580)
	tabWidget.SetCurrentIndex(0)
	tabWidget.SetStyleSheet(`QTabBar::tab {min-width: 75px; background: rgba(255, 255, 255, 1);}
                                        QTabBar::tab:selected {border-bottom: 2px solid #4796f0;}
                                        QLabel{"background: transparent;}
                                        QCheckBox{"background: transparent;}`)
	layout.AddWidget(tabWidget, 1, 0)

	// 工具栏1
	tab_1 := widgets.NewQWidget(tabWidget, core.Qt__Widget)
	tabWidget.AddTab2(tab_1, gui.NewQIcon5("config/static/图标.ico"), "")
	tabWidget.SetTabText(tabWidget.IndexOf(tab_1), "API设定")

	// 工具栏2
	tab_2 := widgets.NewQWidget(tabWidget, core.Qt__Widget)
	tabWidget.AddTab2(tab_2, gui.NewQIcon5("config/static/图标.ico"), "")
	tabWidget.SetTabText(tabWidget.IndexOf(tab_2), "翻译源设定")

	// 同步配置按钮
	getSettinButton := widgets.NewQPushButton(tab_1)
	getSettinButton.SetGeometry2(324, 0, 80, 30)
	getSettinButton.SetStyleSheet("background: rgba(255, 255, 255, 0.4);")
	getSettinButton.SetText("同步设置")

	window.SetLayout(layout)

	return window
}
