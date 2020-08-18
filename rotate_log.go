type LogFileWriter struct {
	path     string
	oldList  *list.List
	maxCount int
	maxSize  int64

	file *os.File
	size int64
}

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func (p *LogFileWriter) shouldRotate(length int) bool {
	if p.maxSize > 0 {
		if p.size+int64(length) > p.maxSize {
			return true
		}
	}
	return false

}
func (p *LogFileWriter) Write(data []byte) (n int, err error) {
	if p.file == nil {
		//logSuffix := time.Now().Format("20060102150405")
		// fn := strings.Join([]string{"be.log.", logSuffix}, "")
		f, err := os.OpenFile(p.path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err == nil {
			p.file = f
			fi, err := p.file.Stat()
			if err != nil {
				return 0, errors.New("open file failed, cause by " + err.Error())
			}
			p.size = fi.Size()
		} else {
			return 0, errors.New("open file failed, cause by " + err.Error())
		}
	}
	n = len(data)
	if p.shouldRotate(n) {
		//rollover
		if err := p.Dorotate(); err != nil {
			return 0, errors.New("rotate file err, cause by " + err.Error())
		}

	}
	n, e := p.file.Write(data)
	if e != nil {
		return 0, errors.New("write file failed ,file is " + p.file.Name() + ", cause by " + e.Error())
	}
	p.size += int64(n)
	return n, e
}

func (p *LogFileWriter) Dorotate() (err error) {
	_ = p.file.Close()
	if p.maxCount > 0 {
		for i := p.maxCount - 1; i > 0; i-- {
			sfn := fmt.Sprintf("%s.%d", p.path, i)
			dfn := fmt.Sprintf("%s.%d", p.path, i+1)
			if Exist(sfn) {
				if Exist(dfn) {
					if err := os.Remove(dfn); err != nil {
						return err
					}
				}
				if err := os.Rename(sfn, dfn); err != nil {
					return err
				}
			}
		}
		dfn := fmt.Sprintf("%s.%d", p.path, 1)
		if Exist(dfn) {
			if err := os.Remove(dfn); err != nil {
				return err
			}
		}
		if err := os.Rename(p.path, dfn); err != nil {
			return err
		}
	}
	p.file, err = os.OpenFile(p.path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	p.size = 0
	return
}

