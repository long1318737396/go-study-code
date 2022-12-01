package main
import "fmt"

//从网络下载文件，所有的文件类型均需实现该接口
type FileDownloader interface {
	//打开网络连接
	openConnection()
	//从给定的网络资源地址下载文件
	downloadFromUrl(downloadUrl string) []byte
	//将下载的文件保存到本地
	saveToLocalDisk(fileLocation string)
	//关闭网络连接
	closeConnection()
}

type ImageFile struct{
	width int32
	height int32
}
func (imageFileInstance ImageFile) display(){
	fmt.Println("显示图片内容")
}
func (imageFileInstance ImageFile) openConnection() {
	fmt.Println("准备下载图片文件，打开网络连接")
}
func (imageFileInstance ImageFile) downloadFromUrl(remoteUrl string) []byte{
	fmt.Println("正在下载图片文件")
	return []byte{}
}
func (imageFileInstance ImageFile) saveToLocalDisk(localFileUrl string) {
	fmt.Println("正在保存图片文件")
}
func (imageFileInstance ImageFile) closeConnection() {
	fmt.Println("下载图片文件成功，关闭网络连接")
}
type AudioFile struct{
	duration int64
}
func (audioFileInstance AudioFile) play(){
	fmt.Println("播放音频文件")
}
func (audioFileInstance AudioFile) openConnection() {
	fmt.Println("准备下载音频文件，打开网络连接")
}
func (audioFileInstance AudioFile) downloadFromUrl(string) (data []byte){
	fmt.Println("正在下载音频文件")
	data=[]byte{}
	return
}
func (audioFileInstance AudioFile) saveToLocalDisk(string) {
	fmt.Println("正在保存音频文件")
}
func (audioFileInstance AudioFile) closeConnection() {
	fmt.Println("下载音频文件成功，关闭网络连接")
}


func main(){
	var fileDownloader FileDownloader
	imageFile:=ImageFile{}
	fileDownloader=imageFile
	fileDownloader.openConnection()
	fileDownloader.downloadFromUrl("https://xxx.xxx.xxx/xxx.jpg")
	fileDownloader.saveToLocalDisk("D:\\xxx.jpg")
	fileDownloader.closeConnection()
	audioFile:=AudioFile{}
	fileDownloader=audioFile
	fileDownloader.openConnection()
	fileDownloader.downloadFromUrl("https://xxx.xxx.xxx/xxx.m4a")
	fileDownloader.saveToLocalDisk("D:\\xxx.m4a")
	fileDownloader.closeConnection()
	
}