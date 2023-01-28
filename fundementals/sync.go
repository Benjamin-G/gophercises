package fundementals

import (
	"fmt"
	"sync"
	"time"
)

func generateThumbnail(wg *sync.WaitGroup, image string, size int) {
	defer wg.Done()
	// thumbnail to be generated
	thumb := fmt.Sprintf("%s@%dx.png", image, size)
	fmt.Println("Generating thumbnail:", thumb)
	// wait for the thumbnail to be ready
	time.Sleep(time.Millisecond * time.Duration(size))
	fmt.Println("Finished generating thumbnail:", thumb)
}

func generateThumbnailV2(image string, size int) error {
	// error if the size is divisible by 5
	if size%5 == 0 {
		return fmt.Errorf("%d is divisible by 5", size)
	}
	// thumbnail to be generated
	thumb := fmt.Sprintf("%s@%dx.png", image, size)
	fmt.Println("Generating thumbnail:", thumb)
	// wait for the thumbnail to be ready
	time.Sleep(time.Millisecond * time.Duration(size))
	fmt.Println("Finished generating thumbnail:", thumb)
	return nil
}

type Builder struct {
	Built bool
	once  sync.Once
}

func (b *Builder) Build() error {
	var err error
	b.once.Do(func() {
		fmt.Print("building...")
		time.Sleep(10 * time.Millisecond)
		fmt.Println("built")
		b.Built = true
		// validate the message
		if !b.Built {
			err = fmt.Errorf("expected builder to be built")
		}
	})
	// return the b.msg and the error variable
	return err
}

type Manager struct {
	quit chan struct{}
	once sync.Once
}

func (m *Manager) Quit() {
	// close the manager's quit channel
	// this will only close the channel once
	m.once.Do(func() {
		fmt.Println("closing quit channel")
		close(m.quit)
	})
}
