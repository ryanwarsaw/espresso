use cursive::Cursive;
use cursive::theme::{Color, PaletteColor, Theme};
use cursive::traits::{Nameable, Resizable, Scrollable};
use cursive::view::ScrollStrategy;
use cursive::views::{EditView, LinearLayout, ResizedView, ScrollView, TextView};

fn main() {
    let mut cursive_root = Cursive::default();
    cursive_root.set_theme(configure_theme(&cursive_root));
    cursive_root.screen_mut().add_fullscreen_layer(create_message_view());
    cursive_root.run();
}

fn create_message_view() -> LinearLayout {
    LinearLayout::vertical()
        .child(TextView::empty()
            .scrollable()
            .show_scrollbars(false)
            .scroll_strategy(ScrollStrategy::StickToBottom)
            .full_screen()
            .with_name("message_history"))
        .child(EditView::new()
            .on_submit(handle_message_submit)
            .full_width()
            .with_name("send_message"))
}

fn handle_message_submit(cursive_root: &mut Cursive, message: &str) {
    cursive_root.call_on_name("message_history", |view: &mut ResizedView<ScrollView<TextView>>| {
        view.get_inner_mut().get_inner_mut().append(message.to_owned() + "\n");
        view.get_inner_mut().set_scroll_strategy(ScrollStrategy::StickToBottom);
    });

    // Reset the send message input once we have rendered the message.
    cursive_root.call_on_name("send_message", |view: &mut ResizedView<EditView>| {
        view.get_inner_mut().set_content("");
    });
}

fn configure_theme(cursive_root: &Cursive) -> Theme {
    let mut theme = cursive_root.current_theme().clone();
    theme.palette[PaletteColor::View] = Color::TerminalDefault;
    theme.palette[PaletteColor::Primary] = Color::TerminalDefault;
    return theme;
}