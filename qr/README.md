##	二维码相关操作方法

###	用法
go get gitee.com/lyhuilin/pkg/qr


####		打印二维码到终端标准输出流

	// 把二维码输出到标准输出流
	if err = qr.PrintQRCode(code); err != nil {
		return err
	}